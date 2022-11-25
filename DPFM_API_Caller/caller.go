package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-orders-creates-rmq-kube/config"
	"data-platform-api-orders-creates-rmq-kube/existence_conf"
	"data-platform-api-orders-creates-rmq-kube/sub_func_complementer"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	configure    *existence_conf.ExistenceConf
	complementer *sub_func_complementer.SubFuncComplementer
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

	confirmor *existence_conf.ExistenceConf,
	complementer *sub_func_complementer.SubFuncComplementer,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:          context.Background(),
		conf:         conf,
		rmq:          rmq,
		configure:    confirmor,
		complementer: complementer,
	}
}

func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *sub_func_complementer.SDC,
	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exconfAllExist := false

	subFuncFin := make(chan error)
	exconfFin := make(chan error)

	// 他PODへ問い合わせ
	wg.Add(1)
	go func() {
		defer wg.Done()
		var e []error
		exconfAllExist, e = c.configure.Conf(input, output, log)
		if len(e) != 0 {
			mtx.Lock()
			errs = append(errs, e...)
			mtx.Unlock()
			exconfFin <- xerrors.New("exconf error")
			return
		}
		exconfFin <- nil
	}()

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, subFuncFin, log, &errs, input, output)
		case "Item":
			// TODO: 実装
			errs = append(errs, xerrors.New("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}

	// 後処理
	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-exconfFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		errs = append(errs, xerrors.New("time out"))
		return errs
	}

	if !exconfAllExist {
		mtx.Lock()
		return nil
	}
	select {
	case e := <-subFuncFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.New("time out"))
		return errs
	}

	return nil
}

func (c *DPFMAPICaller) headerCreate(
	wg *sync.WaitGroup,
	mtx *sync.Mutex,
	errFin chan error,
	log *logger.Logger,
	errs *[]error,
	sdc *dpfm_api_input_reader.SDC,
	ssdc *sub_func_complementer.SDC,
) {
	var err error = nil
	defer wg.Done()
	defer func() {
		errFin <- err
	}()
	sessionID := sdc.RuntimeSessionID
	ctx := context.Background()
	err = c.complementer.ComplementHeader(sdc, ssdc, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
		return
	}

	// data_platform_orders_header_dataの更新
	headerData := ssdc.Message.Header
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "OrdersHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Header Data cannot insert")
		ssdc.SQLUpdateResult = getBoolPtr(false)
		ssdc.SQLUpdateError = "Header Data cannot insert"
		return
	}

	// data_platform_orders_header_partner_dataの更新
	for i := range ssdc.Message.HeaderPartner {
		headerPartnerData := ssdc.Message.HeaderPartner[i]
		res, err = c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerPartnerData, "function": "OrdersHeaderPartner", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return
		}
		res.Success()
	}
	if !checkResult(res) {
		// err = xerrors.New("Header Partner Data cannot insert")
		ssdc.SQLUpdateResult = getBoolPtr(false)
		ssdc.SQLUpdateError = "Header Partner Data cannot insert"
		return
	}

	// data_platform_orders_header_partner_plant_dataの更新
	for i := range ssdc.Message.HeaderPartnerPlant {
		headerPartnerPlantData := ssdc.Message.HeaderPartnerPlant[i]
		res, err = c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerPartnerPlantData, "function": "OrdersHeaderPartnerPlant", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return
		}
		res.Success()
	}
	if !checkResult(res) {
		// err = xerrors.Errorf("Header Partner Plant Data cannot insert")
		ssdc.SQLUpdateResult = getBoolPtr(false)
		ssdc.SQLUpdateError = "Header Partner Plant Data cannot insert"
		return
	}

	ssdc.SQLUpdateResult = getBoolPtr(true)
	return
}

func (c *DPFMAPICaller) itemCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, input *dpfm_api_input_reader.SDC) {
	return
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	_, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := data["result"].(string)
	if !ok {
		return false
	}
	return result == "success"

}

func getBoolPtr(b bool) *bool {
	return &b
}

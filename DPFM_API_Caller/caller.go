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

// type RMQOutputter interface {
// 	Send(sendQueue string, payload map[string]interface{}) error
// }

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	confirmor    *existence_conf.ExistenceConf
	complementer *sub_func_complementer.SubFuncComplementer

	// outputter RMQOutputter
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

	confirmor *existence_conf.ExistenceConf,
	complementer *sub_func_complementer.SubFuncComplementer,
	// outputter RMQOutputter,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:          context.Background(),
		conf:         conf,
		rmq:          rmq,
		confirmor:    confirmor,
		complementer: complementer,
		// outputter:    outputter,
	}
}

func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

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
		exconfAllExist, e = c.confirmor.Conf(input, log)
		if len(e) != 0 {
			mtx.Lock()
			errs = append(errs, e...)
			mtx.Unlock()
			exconfFin <- xerrors.Errorf("exconf error")
			return
		}
		exconfFin <- nil
	}()

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, subFuncFin, log, errs, input)
		case "Item":
			// TODO: 実装
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
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
		errs = append(errs, xerrors.Errorf("time out"))
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
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	log.Info(input)
	return nil
}

func (c *DPFMAPICaller) headerCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, orders *dpfm_api_input_reader.SDC) {
	defer wg.Done()
	err := c.complementer.ComplementHeader(orders, log)
	if err != nil {
		mtx.Lock()
		errs = append(errs, err)
		mtx.Unlock()
		errFin <- xerrors.Errorf("complement error")
	}
	errFin <- nil

	// headerData, err := c.callToHeader("A_BusinessPartner", businessPartner)
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }
	headerData := orders.Orders
	err = c.rmq.Send(c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "OrdersHeader"})

	// headerData := input
	// err = c.rmq.Send(c.conf.RMQ.QueueToSQL()[0], headerData)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(map[string]interface{}{"message": headerData, "function": "OrdersHeader"})
}

// func (c *DPFMAPICaller) callToHeader(input *dpfm_api_input_reader.Input, l *logger.Logger) (*dpfm_api_output_formatter.Header, error) {

// 	data, err := dpfm_api_output_formatter.ConvertToHeader(byteArray, l)
// 	if err != nil {
// 		return nil, xerrors.Errorf("convert error: %w", err)
// 	}
// 	return data, err
// }

func (c *DPFMAPICaller) itemCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, input *dpfm_api_input_reader.SDC) {
	return
}

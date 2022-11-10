package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-orders-creates-rmq-kube/config"
	"data-platform-api-orders-creates-rmq-kube/existence_check"
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

	checker      *existence_check.ExistenceChecker
	complementer *sub_func_complementer.SubFuncComplementer
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

	checker *existence_check.ExistenceChecker,
	complementer *sub_func_complementer.SubFuncComplementer,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:          context.Background(),
		conf:         conf,
		rmq:          rmq,
		checker:      checker,
		complementer: complementer,
	}
}

func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.Input,

	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exMap := make(map[string]bool)
	exconfAllExist := false
	errFin := make(chan error)

	// 他PODへ問い合わせ
	wg.Add(1)
	go func() {
		defer wg.Done()
		var e []error
		exMap, exconfAllExist, e = c.checker.Check(input, log)
		if len(e) != 0 {
			mtx.Lock()
			errs = append(errs, e...)
			mtx.Unlock()
			log.Error(exconfAllExist)
			errFin <- xerrors.Errorf("exconf error")
			return
		}
		if !exconfAllExist {
			errFin <- xerrors.Errorf("exconf not exist")
			return
		}
		log.Info(exMap)
	}()

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, errFin, log, errs, input)
		case "Item":
			// TODO: 実装
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}

	// 後処理
	fin := make(chan struct{})
	go func() {
		wg.Wait()
		fin <- struct{}{}
	}()
	ticker := time.NewTicker(10 * time.Second)

	select {
	case <-fin:
		break
	case e := <-errFin:
		mtx.Lock()
		errs = append(errs, e)
		return errs
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	if !exconfAllExist {
		errs = append(errs, xerrors.Errorf("%v", exMap))
	}
	return nil
}

func (c *DPFMAPICaller) headerCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, input *dpfm_api_input_reader.Input) {
	defer wg.Done()
	err := c.complementer.ComplementHeader(input, log)
	if err != nil {
		mtx.Lock()
		errs = append(errs, err)
		mtx.Unlock()
		errFin <- xerrors.Errorf("complement error")
	}
}

func (c *DPFMAPICaller) itemCreate(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs []error, input *dpfm_api_input_reader.Input) {

	return
}

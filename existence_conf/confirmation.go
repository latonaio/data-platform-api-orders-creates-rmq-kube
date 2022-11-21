package existence_conf

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-orders-creates-rmq-kube/config"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context

	c             *config.Conf
	queueToMapper exconfQueueMapper
	rmq           *rabbitmq.RabbitmqClient
}

func NewExistenceConf(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient) *ExistenceConf {
	return &ExistenceConf{
		ctx:           ctx,
		c:             c,
		queueToMapper: NewExconfQueueMapper(c),
		rmq:           rmq,
	}
}

// Confirm returns existenceMap, allExist, err
func (c *ExistenceConf) Conf(data *dpfm_api_input_reader.SDC, l *logger.Logger) (allExist bool, errs []error) {
	existenceMap := make([]bool, 0, 5)
	wg := sync.WaitGroup{}
	mtx := &sync.Mutex{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		err := c.bpExistenceConf(*data.Orders.Buyer, data, &existenceMap, mtx, l)
		if err != nil {
			mtx.Lock()
			errs = append(errs, err)
			mtx.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err := c.bpExistenceConf(*data.Orders.Seller, data, &existenceMap, mtx, l)
		if errs != nil {
			mtx.Lock()
			errs = append(errs, err)
			mtx.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		err := c.plantExistenceConf(data.Orders.HeaderPartner, data, &existenceMap, mtx, l)
		if errs != nil {
			mtx.Lock()
			errs = append(errs, err)
			mtx.Unlock()
		}
	}()

	wg.Wait()
	if len(errs) != 0 {
		return false, errs
	}

	for _, v := range existenceMap {
		if v {
			continue
		}
		return false, nil
	}
	return true, nil
}

func confKeyExistence(res map[string]interface{}) bool {
	if res == nil {
		return false
	}
	raw, ok := res["ExistenceConf"]
	exist := fmt.Sprintf("%v", raw)
	if ok {
		return strings.ToLower(exist) == "true"
	}

	return false
}

func (c *ExistenceConf) bpExistenceConf(bpID int, data *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) error {
	key := "BusinessPartnerGeneral"
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()

	}()
	b, _ := json.Marshal(data)
	req := BusinessPartnerReq{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return xerrors.Errorf("Unmarshal error: %w", err)
	}

	req.BusinessPartner.BusinessPartner = bpID
	res, err := c.rmq.SessionKeepRequest(nil, c.queueToMapper[key], req)
	if err != nil {
		return xerrors.Errorf("response error: %w", err)
	}
	res.Success()
	exist = confKeyExistence(res.Data())
	log.Info(res.Data())
	return nil
}

func (c *ExistenceConf) plantExistenceConf(headerPartners []dpfm_api_input_reader.HeaderPartner, data *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) error {
	key := "PlantGeneral"
	exist := make([]bool, 0, len(headerPartners))
	defer func() {
		if len(exist) == 0 {
			mtx.Lock()
			*existenceMap = append(*existenceMap, false)
			mtx.Unlock()
			return
		}
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist...)
		mtx.Unlock()
	}()

	b, _ := json.Marshal(data)
	req := PlantReq{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return xerrors.Errorf("Unmarshal error: %w", err)
	}
	wg := sync.WaitGroup{}
	for _, v := range headerPartners {
		wg.Add(1)
		go func(req PlantReq, hp dpfm_api_input_reader.HeaderPartner) {
			defer wg.Done()
			req.Plant.BusinessPartner = *hp.BusinessPartner

			for _, p := range hp.HeaderPartnerPlant {
				plant := p.Plant
				if plant == "" {
					exist = append(exist, true)
					continue
				}
				req.Plant.Plant = plant
				res, err := c.rmq.SessionKeepRequest(nil, c.queueToMapper[key], req)
				if err != nil {
					log.Error(xerrors.Errorf("response error: %w", err))
				}
				res.Success()
				exist = append(exist, confKeyExistence(res.Data()))
				log.Info(res.Data())
			}

		}(req, v)
	}
	wg.Wait()

	return nil
}

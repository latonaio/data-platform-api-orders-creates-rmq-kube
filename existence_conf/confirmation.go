package existence_conf

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-orders-creates-rmq-kube/config"
	"data-platform-api-orders-creates-rmq-kube/database"
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
	db            *database.Mysql
}

func NewExistenceConf(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql) *ExistenceConf {
	return &ExistenceConf{
		ctx:           ctx,
		c:             c,
		queueToMapper: NewExconfQueueMapper(c),
		rmq:           rmq,
		db:            db,
	}
}

// Confirm returns existenceMap, allExist, err
func (c *ExistenceConf) Conf(data *dpfm_api_input_reader.SDC, l *logger.Logger) (allExist bool, errs []error) {
	existenceMap := make(map[int]bool, 2)
	wg := sync.WaitGroup{}
	mtx := &sync.Mutex{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := c.bpExistenceConf(*data.Orders.Buyer, data, existenceMap, mtx, l)
		if err != nil {
			mtx.Lock()
			errs = append(errs, err)
			mtx.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err := c.bpExistenceConf(*data.Orders.Seller, data, existenceMap, mtx, l)
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

func (c *ExistenceConf) bpExistenceConf(bpID int, data *dpfm_api_input_reader.SDC, existenceMap map[int]bool, mtx *sync.Mutex, log *logger.Logger) error {
	key := "BusinessPartner"
	mtx.Lock()
	existenceMap[bpID] = false
	mtx.Unlock()
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
	exist := confKeyExistence(res.Data())
	log.Info(res.Data())
	mtx.Lock()
	existenceMap[bpID] = exist
	mtx.Unlock()
	return nil
}

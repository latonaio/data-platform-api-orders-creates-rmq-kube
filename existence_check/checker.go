package existence_check

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

type ExistenceChecker struct {
	ctx context.Context

	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
	db  *database.Mysql
}

func NewExistenceChecker(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient, db *database.Mysql) *ExistenceChecker {
	return &ExistenceChecker{
		ctx: ctx,
		c:   c,
		rmq: rmq,
		db:  db,
	}
}

// Check returns existenceMap, allExist, err
func (c *ExistenceChecker) Check(data *dpfm_api_input_reader.Input, l *logger.Logger) (existenceMap map[string]bool, allExist bool, err []error) {
	existenceMap = make(map[string]bool, 2)
	wg := sync.WaitGroup{}
	mtx := &sync.Mutex{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		e := c.bpExistenceCheck(*data.Orders.Buyer, data, existenceMap, mtx)
		if e != nil {
			mtx.Lock()
			err = append(err, e)
			mtx.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		e := c.bpExistenceCheck(*data.Orders.Seller, data, existenceMap, mtx)
		if err != nil {
			mtx.Lock()
			err = append(err, e)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	if len(err) != 0 {
		return existenceMap, false, err
	}

	for _, v := range existenceMap {
		if v {
			continue
		}
		return existenceMap, false, nil
	}
	return existenceMap, true, nil
}

func checkKeyExistence(res map[string]interface{}) bool {
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

func (c *ExistenceChecker) bpExistenceCheck(bpID int, data *dpfm_api_input_reader.Input, existenceMap map[string]bool, mtx *sync.Mutex) error {
	key := "BusinessPartner"
	mtx.Lock()
	existenceMap[key] = false
	mtx.Unlock()

	b, _ := json.Marshal(data)

	req := BusinessPartnerReq{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return xerrors.Errorf("Unmarshal error: %w", err)
	}

	req.BusinessPartner.BusinessPartner = bpID
	res, err := c.rmq.SessionKeepRequest(nil, getQueueNameByCheckContent(key), req)
	if err != nil {
		return xerrors.Errorf("response error: %w", err)
	}
	res.Success()
	exist := checkKeyExistence(res.Data())
	mtx.Lock()
	existenceMap[key] = exist
	mtx.Unlock()

	return nil
}

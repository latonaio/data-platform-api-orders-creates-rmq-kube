package existence_conf

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-creates-rmq-kube/DPFM_API_Output_Formatter"
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
func (c *ExistenceConf) Conf(data *dpfm_api_input_reader.SDC, ssdc *dpfm_api_output_formatter.SDC, l *logger.Logger) (allExist bool, errs []error) {
	var resMsg string
	existenceMap := make([]bool, 0, 5)
	wg := sync.WaitGroup{}
	mtx := &sync.Mutex{}

	wg.Add(3)
	go c.bpExistenceConf(*data.Header.Buyer, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
	go c.bpExistenceConf(*data.Header.Seller, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)
	go c.plantExistenceConf(data.Header.HeaderPartner, data, &existenceMap, &resMsg, &errs, mtx, &wg, l)

	wg.Wait()

	if len(errs) != 0 {
		return false, errs
	}

	ssdc.ExconfResult = getBoolPtr(true)
	for _, v := range existenceMap {
		if v {
			continue
		}
		ssdc.ExconfResult = getBoolPtr(false)
		ssdc.ExconfError = resMsg
		return false, nil
	}
	return true, nil
}

func (c *ExistenceConf) bpExistenceConf(bpID int, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	res, err := c.bpExistenceConfRequest(bpID, input, existenceMap, mtx, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
	}
	if res != "" {
		*exconfErrMsg = res
	}
}

func (c *ExistenceConf) plantExistenceConf(headerPartners []dpfm_api_input_reader.HeaderPartner, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0
	for _, hp := range headerPartners {
		if len(hp.HeaderPartnerPlant) == 0 {
			*exconfErrMsg = plantNotExist()
		}
		for _, p := range hp.HeaderPartnerPlant {
			wg2.Add(1)
			exReqTimes++
			go func(plant string, bpID int) {
				res, err := c.plantExistenceConfRequest(plant, bpID, input, existenceMap, mtx, log)
				if err != nil {
					mtx.Lock()
					*errs = append(*errs, err)
					mtx.Unlock()
				}
				if res != "" {
					*exconfErrMsg = res
				}
				wg2.Done()
			}(p.Plant, *hp.BusinessPartner)
		}
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) bpExistenceConfRequest(bpID int, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	key := "BusinessPartnerGeneral"
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[BusinessPartnerReq](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.BusinessPartner.BusinessPartner = bpID

	exist, err = c.exconfRequest(req, key, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}

	return "", nil
}

func (c *ExistenceConf) plantExistenceConfRequest(plant string, bpID int, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	key := "PlantGeneral"
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Plant":           plant,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[PlantReq](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.Plant.Plant = plant
	req.Plant.BusinessPartner = bpID

	exist, err = c.exconfRequest(req, key, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func jsonTypeConversion[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
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
func (c *ExistenceConf) exconfRequest(req interface{}, key string, log *logger.Logger) (bool, error) {
	res, err := c.rmq.SessionKeepRequest(nil, c.queueToMapper[key], req)
	if err != nil {
		return false, xerrors.Errorf("response error: %w", err)
	}
	res.Success()
	exist := confKeyExistence(res.Data())
	log.Info(res.Data())
	return exist, nil
}

type result struct {
	keys map[string]interface{}
}

func newResult(keys map[string]interface{}) *result {
	return &result{
		keys: keys,
	}
}

func (r *result) fail() string {
	txt := ""
	for k, v := range r.keys {
		txt = fmt.Sprintf("%s%s:%v, ", k, v)
	}
	txt = fmt.Sprintf("%s does not exist", txt)
	return txt
}
func plantNotExist() string {
	return "plant data does not exist."
}

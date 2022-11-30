package sub_func_complementer

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-orders-creates-rmq-kube/config"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SubFuncComplementer struct {
	ctx context.Context
	c   *config.Conf
	rmq *rabbitmq.RabbitmqClient
}

func NewSubFuncComplementer(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient) *SubFuncComplementer {
	return &SubFuncComplementer{
		ctx: ctx,
		c:   c,
		rmq: rmq,
	}
}

func (c *SubFuncComplementer) ComplementHeader(data *dpfm_api_input_reader.SDC, ssdc *SDC, l *logger.Logger) error {
	s := &SDC{}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Headers"], data)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}
	ssdc.Message.Header = s.Message.Header
	ssdc.Message.HeaderPartner = s.Message.HeaderPartner
	ssdc.Message.HeaderPartnerPlant = s.Message.HeaderPartnerPlant

	ssdc.SubfuncResult = getBoolPtr(true)
	if s.SubfuncResult == nil || !*s.SubfuncResult {
		ssdc.SubfuncResult = getBoolPtr(false)
		ssdc.SubfuncError = s.SubfuncError
		return xerrors.New(ssdc.SubfuncError)
	}

	return err
}

func (c *SubFuncComplementer) ComplementItem(data *dpfm_api_input_reader.SDC, ssdc *SDC, l *logger.Logger) error {
	s := &SDC{}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Items"], data)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}
	ssdc.Message.Item = s.Message.Item

	ssdc.SubfuncResult = getBoolPtr(true)
	if s.SubfuncResult == nil || !*s.SubfuncResult {
		ssdc.SubfuncResult = getBoolPtr(false)
		ssdc.SubfuncError = s.SubfuncError
		return xerrors.New(ssdc.SubfuncError)
	}

	return err
}

func getBoolPtr(b bool) *bool {
	return &b
}

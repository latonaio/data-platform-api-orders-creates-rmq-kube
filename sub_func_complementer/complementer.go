package sub_func_complementer

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-creates-rmq-kube/DPFM_API_Output_Formatter"
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

func (c *SubFuncComplementer) ComplementHeader(input *dpfm_api_input_reader.SDC, output *dpfm_api_output_formatter.SDC, outputMsg *dpfm_api_output_formatter.CreatesMessage, l *logger.Logger) error {
	s := &dpfm_api_output_formatter.SDC{}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Headers"], input)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(s.Message)
	msg := &dpfm_api_output_formatter.CreatesMessage{}
	err = json.Unmarshal(b, msg)
	if err != nil {
		return err
	}
	outputMsg.HeaderCreates = msg.HeaderCreates
	outputMsg.HeaderPartner = msg.HeaderPartner
	outputMsg.HeaderPartnerPlant = msg.HeaderPartnerPlant

	output.SubfuncResult = getBoolPtr(true)
	if s.SubfuncResult == nil || !*s.SubfuncResult {
		output.SubfuncResult = getBoolPtr(false)
		output.SubfuncError = s.SubfuncError
		return xerrors.New(output.SubfuncError)
	}

	return err
}

func (c *SubFuncComplementer) ComplementItem(input *dpfm_api_input_reader.SDC, output *dpfm_api_output_formatter.SDC, outputMsg *dpfm_api_output_formatter.CreatesMessage, l *logger.Logger) error {
	s := &dpfm_api_output_formatter.SDC{}
	res, err := c.rmq.SessionKeepRequest(nil, c.c.RMQ.QueueToSubFunc()["Items"], input)
	if err != nil {
		return err
	}
	res.Success()

	err = json.Unmarshal(res.Raw(), s)
	if err != nil {
		return err
	}
	b, _ := json.Marshal(s.Message)
	msg := &dpfm_api_output_formatter.CreatesMessage{}
	err = json.Unmarshal(b, msg)
	if err != nil {
		return err
	}
	outputMsg.Item = msg.Item

	output.SubfuncResult = getBoolPtr(true)
	if s.SubfuncResult == nil || !*s.SubfuncResult {
		output.SubfuncResult = getBoolPtr(false)
		output.SubfuncError = s.SubfuncError
		return xerrors.New(output.SubfuncError)
	}

	return err
}

func getBoolPtr(b bool) *bool {
	return &b
}

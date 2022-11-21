package main

import (
	"context"
	dpfm_api_caller "data-platform-api-orders-creates-rmq-kube/DPFM_API_Caller"
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-orders-creates-rmq-kube/config"
	"data-platform-api-orders-creates-rmq-kube/existence_conf"
	"data-platform-api-orders-creates-rmq-kube/handlers"
	"data-platform-api-orders-creates-rmq-kube/sub_func_complementer"
	"encoding/json"
	"fmt"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.SessionControlQueue(), conf.RMQ.QueueToSQL(), 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	confirmor := existence_conf.NewExistenceConf(ctx, conf, rmq)
	complementer := sub_func_complementer.NewSubFuncComplementer(ctx, conf, rmq)
	caller := dpfm_api_caller.NewDPFMAPICaller(conf, rmq, confirmor, complementer)

	for msg := range iter {
		start := time.Now()
		err = callProcess(rmq, caller, conf, msg)
		if err != nil {
			msg.Fail()
			handlers.ErrorHandler(err, rmq, conf, l)
			continue
		}
		msg.Success()
		l.Info("process time %v\n", time.Since(start).Milliseconds())
	}
}
func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func callProcess(rmq *rabbitmq.RabbitmqClient, caller *dpfm_api_caller.DPFMAPICaller, conf *config.Conf, msg rabbitmq.RabbitmqMessage) (err error) {
	l := logger.NewLogger()
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("error occurred: %w", e)
			l.Error(err)
			return
		}
	}()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
	var input dpfm_api_input_reader.SDC
	err = json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		l.Error(err)
		return
	}

	accepter := getAccepter(&input)

	errs := caller.AsyncOrderCreates(accepter, &input, l)
	if len(errs) != 0 {
		for _, err := range errs {
			l.Error(err)
		}
		return errs[0]
	}
	rmq.Send(conf.RMQ.QueueToResponse(), input)

	return nil
}

func getAccepter(input *dpfm_api_input_reader.SDC) []string {
	accepter := input.Accepter
	if len(input.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"Header", "Item",
		}
	}
	return accepter
}

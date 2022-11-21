package handlers

import (
	"data-platform-api-orders-creates-rmq-kube/config"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func ErrorHandler(err error, rmq *rabbitmq.RabbitmqClient, conf *config.Conf, l *logger.Logger) {
	l.Error(err)

	err = rmq.Send(conf.RMQ.QueueToResponse(), &ErrorResponse{
		ResponseType: typeError,
		Name:         InternalServerError,
		Message:      fmt.Sprintf("%s", err),
	})
	if err != nil {
		l.Error(err)
	}
}

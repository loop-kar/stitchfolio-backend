package log

import (
	"context"

	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/util"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

// TODO: look into go init method
var logger *zap.Logger

func InitLogger(newRelic *newrelic.Application) {

	//Shud we have seperate logger for taskrunner?
	if logger == nil {
		logger = ProvideZap(newRelic)
	}

}

func Get() *zap.Logger {
	return logger
}

func FromCtx(ctx *context.Context) *zap.Logger {

	//If the context already has a zapLogger , return it
	//If not return the original singleton logger we created on AppStart
	val := util.ReadValueFromContext(ctx, constants.LOGGER)
	if l, ok := val.(*zap.Logger); ok {
		return l
	} else {
		return logger
	}
}

func WithCtx(ctx *context.Context, log *zap.Logger) {

	//Check if logger is already present in context
	//If already present then return
	val := util.ReadValueFromContext(ctx, constants.LOGGER)
	if _, ok := val.(*zap.Logger); ok {
		return
	}

	//Feed the logger into context
	util.SetValueToContext(ctx, constants.LOGGER, log)
}

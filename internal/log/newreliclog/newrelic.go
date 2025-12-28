package newreliclog

import (
	"fmt"
	"os"

	"github.com/imkarthi24/sf-backend/internal/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var newRelicApp *newrelic.Application

func ProvideNewRelic(appConfig config.AppConfig) *newrelic.Application {

	if newRelicApp != nil {
		return newRelicApp
	}

	fmt.Println("Initializing Newrelic")

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appConfig.Server.AppName),
		newrelic.ConfigLicense(appConfig.NewRelic.License),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if nil != err {
		fmt.Printf("New Relic initialization failed: %v\n", err)
		os.Exit(1)
	}

	return app
}

func Get() *newrelic.Application {
	return newRelicApp
}

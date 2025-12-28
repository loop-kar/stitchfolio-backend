package config_test

import (
	"os"
	"testing"

	"github.com/imkarthi24/sf-backend/pkg/config"
	"github.com/stretchr/testify/suite"
)

type configSuite struct {
	suite.Suite
	tc TestConfig
}

type ServerConfig struct {
	AppName  string `mapstructure:"appName"`
	LogLevel string `mapstructure:"logLevel"`
	Port     int    `mapstructure:"port"`
}

type TestConfig struct {
	Server ServerConfig `mapstructure:"server"`
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(configSuite))
}
func (c *configSuite) SetupTest() {
	os.Clearenv()
	configReader, err := os.Open("test.yaml")
	c.Require().NoError(err)

	binding := map[string]string{
		"server.logLevel": "LOG_LEVEL",
		"server.port":     "APP_PORT",
	}

	err = os.Setenv("LOG_LEVEL", "info")
	c.Require().NoError(err)

	err = os.Setenv("APP_PORT", "9090")
	c.Require().NoError(err)

	err = config.LoadConfig(configReader, binding, &c.tc)
	c.Require().NoError(err)

}

func (c *configSuite) TestLoadConfig() {
	c.Assert().Equal("SkillFort-backend", c.tc.Server.AppName)
	c.Assert().Equal("info", c.tc.Server.LogLevel)
	c.Assert().Equal(9090, c.tc.Server.Port)

}

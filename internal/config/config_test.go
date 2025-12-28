package config_test

import (
	"os"
	"testing"

	"github.com/imkarthi24/sf-backend/internal/config"
	"github.com/stretchr/testify/suite"
)

type appConfigSuite struct {
	suite.Suite
	config config.AppConfig
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(appConfigSuite))
}
func (c *appConfigSuite) SetupTest() {
	os.Clearenv()

	configReader, err := os.Open("../../config/prod.yaml")

	c.Require().NoError(err)

	envVars := map[string]string{
		"LOG_LEVEL":   "info",
		"APP_PORT":    "9000",
		"DB_HOST":     "host",
		"DB_NAME":     "SkillFort-db",
		"DB_PORT":     "5432",
		"DB_USER":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_SCHEMA":   "ref",
	}

	for key, val := range envVars {
		err = os.Setenv(key, val)
		c.Require().NoError(err)
	}

	c.config, err = config.LoadConfig(configReader)
	c.Require().NoError(err)
}

func (c *appConfigSuite) TestServerConfig() {
	srvConfig := c.config.Server

	c.Require().NotNil(srvConfig, "Expected Server Config to be not null")
	c.Require().Equal("SkillFort-backend", srvConfig.AppName)
	c.Require().Equal("info", srvConfig.LogLevel)
	c.Require().Equal("localhost", srvConfig.Host)
	c.Require().Equal(9000, srvConfig.Port)
}

func (c *appConfigSuite) TestDatabaseConfig() {
	databaseConfig := c.config.Database

	c.Require().NotNil(databaseConfig, "Expected Database Config to be not null")
	c.Require().Equal("SkillFort-db", databaseConfig.DBName)
	c.Require().Equal("host", databaseConfig.Host)
	c.Require().Equal("postgres", databaseConfig.Password)
	c.Require().Equal("postgres", databaseConfig.Username)
	c.Require().Equal(5432, databaseConfig.Port)
	c.Require().Equal("ref", databaseConfig.Schema)

}

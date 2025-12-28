package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicMiddleWare(app *newrelic.Application) gin.HandlerFunc {
	return nrgin.Middleware(app)
}

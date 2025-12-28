package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct{}

func ProvideHealthHandler() Health {
	return Health{}
}

// Health godoc
//
//	@Summary		Health endpoint
//	@Description	get string by ID
//	@Tags			Health
//	@Success		200
//	@Failure		400
//	@Router			/health [get]
func (h Health) Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Healthy")
}

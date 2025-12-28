package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/log"
	"go.uber.org/zap"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		logger := log.Get()

		//Creating and setting correlation id before the req life cycle
		//When keys are copied over from gincontext to appcontext at start of every handler , the
		//correlationId is passed over to appContext
		correlationID := uuid.NewString()

		logger = logger.With(zap.String(constants.CORRELATION_ID, correlationID))
		ctx.Set(constants.CORRELATION_ID, correlationID)
		ctx.Set(constants.LOGGER, logger)

		requestBody := log.HandleRequestBody(ctx.Request)
		respWriter := log.HandleResponseBody(&ctx.Writer)
		ctx.Next()

		logMessage := log.FormatRequestAndResponse(
			ctx.Writer,
			ctx.Request,
			respWriter.Format(),
			correlationID,
			requestBody,
		)

		if logMessage != "" {
			if isSuccessStatusCode(ctx.Writer.Status()) {
				logger.Info(logMessage)
			} else {
				logger.Error(logMessage)
			}
		}
	}
}

func isSuccessStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent:
		return true
	default:
		return false
	}
}

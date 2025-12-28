package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/db"
)

func RequestParser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var paginationModel db.Pagination
		if ctx.Request.Method == "GET" {

			// Read the QueryParam and set the pagination info
			// for the reques
			if ctx.BindQuery(&paginationModel) == nil {
				ctx.Set(constants.PAGINATION_KEY, &paginationModel)
			}

			filter := ctx.Query(constants.FILTER_QUERY_PARAM)
			ctx.Set(constants.FILTER_KEY, filter)

		}
		ctx.Next()
	}
}

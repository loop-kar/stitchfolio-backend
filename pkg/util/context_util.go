package util

import (
	"context"

	"github.com/gin-gonic/gin"
)

func NewContextWithValue(ctx *context.Context, key string, value any) context.Context {

	ginCtx, ok := (*ctx).(*gin.Context)
	if ok {
		ginCtx.Set(key, value)
		return ginCtx
	}

	return context.WithValue(*ctx, key, value)
}

func SetValueToContext(ctx *context.Context, key string, value any) {

	ginCtx, ok := (*ctx).(*gin.Context)
	if ok {
		ginCtx.Set(key, value)
	} else {
		*ctx = context.WithValue(*ctx, key, value)
	}

}

func ReadValueFromContext(ctx *context.Context, key string) interface{} {

	ginCtx, ok := (*ctx).(*gin.Context)

	// if ginCtx == nil {
	// 	return nil
	// }
	//is the if block needed...need understanding

	if ok {
		return ginCtx.Value(key)
	}

	return (*ctx).Value(key)
}

func CopyContextFromGin(ginCtx *gin.Context) context.Context {
	appCtx := ginCtx.Request.Context()

	for k, v := range ginCtx.Keys {
		appCtx = context.WithValue(appCtx, k, v)
	}
	return appCtx
}

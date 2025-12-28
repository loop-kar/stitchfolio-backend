package util

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestNewContextWithValue(t *testing.T) {
	// Test case 1: Regular context
	ctx := context.Background()
	newCtx := NewContextWithValue(&ctx, "testKey", "testValue")
	value := newCtx.Value("testKey")
	require.Equal(t, "testValue", value)

	// Test case 2: Gin context
	gin.SetMode(gin.TestMode)
	ginCtx := &gin.Context{}
	ctx = context.Background()
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	newCtx = NewContextWithValue(&ctx, "testKey", "testValue")
	value = newCtx.Value("testKey")
	require.Equal(t, "testValue", value)
}

func TestSetValueToContext(t *testing.T) {
	// Test case 1: Regular context
	ctx := context.Background()
	SetValueToContext(&ctx, "testKey", "testValue")
	value := ctx.Value("testKey")
	require.Equal(t, "testValue", value)

	// Test case 2: Gin context
	gin.SetMode(gin.TestMode)
	ginCtx := &gin.Context{}
	ctx = context.Background()
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	SetValueToContext(&ctx, "testKey", "testValue")
	value = ctx.Value("testKey")
	require.Equal(t, "testValue", value)
}

func TestReadValueFromContext(t *testing.T) {
	// Test case 1: Regular context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "testKey", "testValue")
	value := ReadValueFromContext(&ctx, "testKey")
	require.Equal(t, "testValue", value)

	// Test case 2: Gin context
	gin.SetMode(gin.TestMode)
	ginCtx := &gin.Context{}
	ctx = context.Background()
	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	ginCtx.Set("testKey", "testValue")
	value = ReadValueFromContext(&ctx, "testKey")
	require.Equal(t, "testValue", value)

	// Test case 3: Non-existent key
	value = ReadValueFromContext(&ctx, "nonExistentKey")
	require.Nil(t, value)
}

func TestCopyContextFromGin(t *testing.T) {
	// Test case 1: Copy context with values
	gin.SetMode(gin.TestMode)
	ginCtx := &gin.Context{}
	ginCtx.Set("key1", "value1")
	ginCtx.Set("key2", "value2")

	appCtx := CopyContextFromGin(ginCtx)
	require.Equal(t, "value1", appCtx.Value("key1"))
	require.Equal(t, "value2", appCtx.Value("key2"))

	// Test case 2: Copy empty context
	ginCtx = &gin.Context{}
	appCtx = CopyContextFromGin(ginCtx)
	require.NotNil(t, appCtx)
}

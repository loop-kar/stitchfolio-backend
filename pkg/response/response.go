package response

import (
	"bytes"
	"context"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"github.com/imkarthi24/sf-backend/pkg/errs"
	"github.com/imkarthi24/sf-backend/pkg/util"
	"gorm.io/gorm"
)

type ResponseType string

const (
	SUCCESS ResponseType = "Success"
	FAILURE ResponseType = "Failure"
	WARNING ResponseType = "Warning"
)

type Response struct {
	Error      *errs.XError   `json:"error,omitempty"`
	Messsage   string         `json:"message"`
	Type       ResponseType   `json:"type"`
	Pagination *db.Pagination `json:"pageMetaData,omitempty"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
	*Response
}

type FileResponse struct {
	FileName string `json:"fileName,omitempty"`
	Content  string `json:"content,omitempty"`
}

func (*Response) DefaultFailureResponse(err *errs.XError) *Response {
	return &Response{
		Messsage: "Error Processing Request",
		Type:     FAILURE,
		Error:    err,
	}
}

func (*Response) FailureResponse(err *errs.XError, message string) *Response {
	return &Response{
		Messsage: message,
		Type:     FAILURE,
		Error:    err,
	}
}

func (*Response) DefaultSuccessResponse() *Response {
	return &Response{
		Messsage: "Request Processed Succesfully",
		Type:     SUCCESS,
	}
}

func (*Response) SuccessResponse(message string) *Response {
	return &Response{
		Messsage: message,
		Type:     SUCCESS,
	}
}

func (*DataResponse) DefaultFailureResponse(err *errs.XError) *DataResponse {
	resp := new(Response)
	return &DataResponse{
		Data:     nil,
		Response: resp.DefaultFailureResponse(err),
	}
}

func (*DataResponse) DefaultSuccessResponse(data interface{}) *DataResponse {
	resp := new(Response)
	return &DataResponse{
		Data:     data,
		Response: resp.DefaultSuccessResponse(),
	}
}

func (*DataResponse) SuccessResponse(data interface{}, message string) *DataResponse {
	resp := new(Response)
	return &DataResponse{
		Data:     data,
		Response: resp.SuccessResponse(message),
	}
}

func (resp *DataResponse) FormatAndSend(appCtx *context.Context, ctx *gin.Context, statusCode int) {

	// if the error has a status code use that
	var code = statusCode
	if resp.Error != nil && resp.Error.Code != 0 {
		code = resp.Error.Code
	}
	disposeTransaction(appCtx, resp.Error)
	resp.setPaginationMetadata(ctx)
	ctx.JSON(code, resp)
}

func (resp *Response) FormatAndSend(appCtx *context.Context, ctx *gin.Context, statusCode int) {

	// if the error has a status code use that
	var code = statusCode
	if resp.Error != nil && resp.Error.Code != 0 {
		code = resp.Error.Code
	}

	disposeTransaction(appCtx, resp.Error)
	resp.setPaginationMetadata(ctx)
	ctx.JSON(code, resp)
}

func (resp *FileResponse) SendFile(appCtx *context.Context, ctx *gin.Context, fileName string, buff bytes.Buffer) {
	disposeTransaction(appCtx, nil)

	bytes := base64.StdEncoding.EncodeToString(buff.Bytes())
	ctx.Header("Content-Description", "File Transfer")
	response := FileResponse{
		FileName: fileName,
		Content:  bytes,
	}
	ctx.JSON(http.StatusOK, response)

}

func (resp *Response) setPaginationMetadata(ctx *gin.Context) {
	pagination := ctx.Value(constants.PAGINATION_KEY)
	if pagination == nil {
		return
	}
	resp.Pagination = pagination.(*db.Pagination)
}

func (resp *DataResponse) setPaginationMetadata(ctx *gin.Context) {
	pagination := ctx.Value(constants.PAGINATION_KEY)
	if pagination == nil {
		return
	}
	resp.Pagination = pagination.(*db.Pagination)
}

func disposeTransaction(ctx *context.Context, err *errs.XError) {
	transaction := getTransactionFromContext(ctx)

	if transaction == nil {
		return
	}

	if err != nil {
		transaction.Rollback()
		return
	}

	transaction.Commit()

}

func getTransactionFromContext(ctx *context.Context) *gorm.DB {
	transaction := util.ReadValueFromContext(ctx, constants.TRANSACTION_KEY)
	if transaction == nil {
		return nil
	}
	return transaction.(*gorm.DB)
}

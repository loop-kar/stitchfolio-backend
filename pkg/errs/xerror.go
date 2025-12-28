package errs

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type XError struct {
	Type    ErrorType     `json:"errorType,omitempty"`
	Message string        `json:"errorCode,omitempty"`
	Err     error         `json:"error,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
	Code    int           `json:"-"`
}

func NewXError(kind ErrorType, msg string, err error) *XError {
	constMsg := strcase.ToScreamingSnake(msg)
	return &XError{
		Type:    kind,
		Message: constMsg,
		Err:     err,
	}
}

func NewParamsXError(kind ErrorType, err error, msg string, params ...interface{}) *XError {
	constMsg := strcase.ToScreamingSnake(msg)
	return &XError{
		Type:    kind,
		Message: constMsg,
		Err:     err,
		Params:  params,
	}
}

func (e *XError) Error() string {
	return fmt.Sprintf("Type : %s | Message: %s | Error: %s", e.Type, e.getFormattedMessage(), e.getError())
}

func Wrap(err error, msg ...string) *XError {
	return &XError{
		Message: strings.Join(msg, " || "),
		Err:     err,
		Type:    OTHER,
	}
}

func (e *XError) SetCode(code int) *XError {
	e.Code = code
	return e
}

func (e *XError) getFormattedMessage() string {
	if len(e.Params) > 0 {
		for i, val := range e.Params {
			e.Message = strings.Replace(e.Message, fmt.Sprintf("{%d}", i), fmt.Sprintf("%v", val), 1)
		}
	}
	return e.Message
}

func (e *XError) getError() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

package task

import "github.com/imkarthi24/sf-backend/pkg/errs"

type ITaskRunner interface {
	Process() *errs.XError
	Run() *errs.XError
	HandleError(*errs.XError)
}

type IBaseTask interface {
	BeforeFetchEntitySet() *errs.XError

	FetchEntitySet() (bool, []TaskResponse, *errs.XError)

	BeforeProcessEntitySet([]TaskResponse) *errs.XError

	ProcessEntitySet([]TaskResponse) (bool, *errs.XError)

	AfterProcessEntitySet(isProcessSuccess bool) *errs.XError

	Dispose(*errs.XError)

	Run() func() *errs.XError
}

type TaskResponse interface{}

package task

import (
	"context"

	"github.com/imkarthi24/sf-backend/pkg/errs"
)

// Implements IBaseTask interface
type BaseTask struct {
	Param *BaseTaskParam
	Ctx   *context.Context
}

func (BaseTask) AfterProcessEntitySet(isProcessSuccess bool) *errs.XError {
	//log Skipping
	return nil
}

func (BaseTask) BeforeFetchEntitySet() *errs.XError {
	//log Skipping
	return nil
}

func (BaseTask) BeforeProcessEntitySet([]TaskResponse) *errs.XError {
	//log Skipping
	return nil
}

func (t BaseTask) Dispose(err *errs.XError) {
	//Commit the transaction
	disposeTransaction(t.Ctx, err)

}

func (BaseTask) FetchEntitySet() (bool, []TaskResponse, *errs.XError) {
	//needs to be implemented by the underlying child
	panic("unimplemented")
}

func (BaseTask) ProcessEntitySet([]TaskResponse) (bool, *errs.XError) {
	//needs to be implemented by the underlying child
	panic("unimplemented")
}

func (BaseTask) Run() func() *errs.XError {
	return nil
}

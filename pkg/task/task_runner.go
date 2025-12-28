package task

import (
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

type taskRunner struct {
	Task  IBaseTask
	Param BaseTaskParam
}

func provideTaskRunner(task IBaseTask, param BaseTaskParam) ITaskRunner {
	return &taskRunner{
		Task:  task,
		Param: param,
	}
}

func (runner *taskRunner) Process() *errs.XError {

	//Before Process
	err := runner.Task.BeforeFetchEntitySet()
	if err != nil {
		return err
	}

	isFetchComplete := false
	for !isFetchComplete {

		var records []TaskResponse
		var err *errs.XError

		//Fetch Entity
		isFetchComplete, records, err = runner.Task.FetchEntitySet()
		if err != nil {
			return err
		}

		//Before Process
		err = runner.Task.BeforeProcessEntitySet(records)
		if err != nil {
			return err
		}

		//Process
		isSuccess, err := runner.Task.ProcessEntitySet(records)
		if err != nil {
			return err
		}

		//After Process
		err = runner.Task.AfterProcessEntitySet(isSuccess)
		if err != nil {
			return err
		}

		if !isSuccess && runner.Param.AbortProceesExecutionOnFailure {
			break
		}

	}

	runner.Task.Dispose(nil)

	return nil

}

func (runner *taskRunner) Run() *errs.XError {

	runner.Task.Run()()
	runner.Task.Dispose(nil)

	return nil

}

func (t *taskRunner) HandleError(err *errs.XError) {
	t.Task.Dispose(err)
}

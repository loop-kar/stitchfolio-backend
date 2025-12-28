package task

// Provides an implementation of JobRunner which allows to run a job
type JobRunner interface {
	// starts a job instantly
	CreateAdHocJob(runImmediate bool)
	RunImmediately()
}

type jobRunner struct {
	tasker ITaskRunner
}

func ProvideJobRunner(task IBaseTask, param BaseTaskParam) JobRunner {
	return &jobRunner{
		tasker: provideTaskRunner(task, param),
	}
}

func (j *jobRunner) CreateAdHocJob(runImmediate bool) {
	if runImmediate {
		j.tasker.Process()
	}
}

func (j *jobRunner) RunImmediately() {
	j.tasker.Run()
}

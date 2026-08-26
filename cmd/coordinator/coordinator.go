package coordinator



type TaskState int

const (
	TaskStateIdle = iota
	TaskStateInProgress
	TaskStateCompleted
)

type TaskMetaData struct{
	State TaskState
	StartTime time.Now()
	TaskID string
	File string
}



type Coordinator struct{
	mu sync.Mutex
	nMap int
	nReduce int
	mapTasks []TaskMetaData
	reduceTasks []TaskMetaData
}


func NewCoordinator(files string , nReduce int) *Coordinator{
	return &Cordinator
}

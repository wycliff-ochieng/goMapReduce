package rpc


import (
	"hash/fnv"
)


//defines whether a task needs to map, reduce, wait or terminate
type TaskType int

const (
	TaskTypeMap TaskType = iota
	TasktypeReduce
	TaskTypeWait
	TaskTypeExit
)


// TaskRequest - rpc payload sent by a worker asking for worker
type TaskRequest struct {
	WorkerID string `json:"WorkerID"`
}

// TaskResponse - rpc payload returned by coordinator with assigned tasks
type TaskResponse struct{
	Type TaskType
	TaskID string  // task identifier - uuid.b6443df7-d18a-4d4f-9ec8-42b1bf104699
	NReduce int // no.of reduce buckets
	NMap int //no.of map buckets
	Input string //file path for map tasks
	Files []string //bin/
}

//ReportResponse - confirms task completion receipt
type ReportResponse struct{
	Type TaskType
	TaskID string
	WorkerID string
	intermediateFiles []string
}

//KeyValue - Represents fundamental intermediate data type emitted 
type KeyValue struct{
	Accepted bool
}

//InHash - distributes keys evenly across R reduce() partitions

func inHash(key string) int{
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum & 0x7fffff)
}

// CoordinatorSock defines the unix domain socket path for local RPC communication
func CoordinatorSock() string{
	return "/tmp/mr-socket-"+ strconv.Itoa(1000)
 }

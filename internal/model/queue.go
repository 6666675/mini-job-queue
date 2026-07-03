package model

import (
	"sync"
	"time"
)

type Queue struct {
	jobs   map[int64]*Job
	queue  chan int64
	nextId int64
	mutex  sync.Mutex
}

// QueueInit 初始化任务队列
func QueueInit() *Queue {
	return &Queue{jobs: make(map[int64]*Job), queue: make(chan int64, 1024), nextId: 0}
}

// Push 创建队列并将任务入队
func (q *Queue) Push(name string, payload []byte) *Job {
	job := new(Job)
	q.mutex.Lock()
	*job = Job{
		Id:         q.nextId,
		Name:       name,
		Payload:    payload,
		CreateTime: time.Now(),
		Status:     Waiting,
	}
	q.jobs[q.nextId] = job
	q.queue <- q.nextId
	q.nextId++
	q.mutex.Unlock()
	return job
}

// Pop 退出队列并运行任务
func (q *Queue) Pop() *Job {
	id := <-q.queue
	q.jobs[id].Status = Running
	q.jobs[id].StartTime = new(time.Now())
	return q.jobs[id]
}
func (q *Queue) Successed(id int64) {
	q.jobs[id].Status = Success
	q.jobs[id].EndTime = new(time.Now())
}
func (q *Queue) Failed(id int64) {
	q.jobs[id].Status = Failed
	q.jobs[id].EndTime = new(time.Now())
}
func (q *Queue) Get(id int64) *Job {
	return q.jobs[id]
}
func (q *Queue) Requeue(id int64) {
	q.jobs[id].Status = Waiting
	q.queue <- id
}

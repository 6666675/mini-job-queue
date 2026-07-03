package service

import (
	"mini-job-queue/internal/model"
	"time"
)

func Worker(q *model.Queue, r *model.Registry, limits map[string]chan struct{}) {
	for {
		job := q.Pop()
		h, ok := r.Get(job.Name)
		if ok != true {
			q.Failed(job.Id)
			continue
		}
		//限制使用channel限制协程数量
		if limits[job.Name] == nil {
			q.Failed(job.Id)
			continue
		}
		select {
		case limits[job.Name] <- struct{}{}:
			{
				go func(job *model.Job, limits map[string]chan struct{}, h model.Handler) {
					defer func() { <-limits[job.Name] }()
					err := h.Run(job.Payload)
					if err != nil {
						q.Failed(job.Id)
						return
					}
					q.Successed(job.Id)
				}(job, limits, h)
			}
		default:
			{
				//如果有超出指定协程数量的handler重新进入队列
				time.Sleep(time.Millisecond * 100)
				q.Requeue(job.Id)
			}
		}
	}
}

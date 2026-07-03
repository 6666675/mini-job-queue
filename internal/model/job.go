package model

import (
	"time"
)

const (
	Waiting int8 = iota
	Running
	Success
	Failed
)

type Job struct {
	Status     int8
	Id         int64
	Name       string
	Payload    []byte
	CreateTime time.Time
	StartTime  *time.Time
	EndTime    *time.Time
}

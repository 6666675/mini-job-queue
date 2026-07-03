package main

import (
	"fmt"
	"mini-job-queue/internal/model"
	"mini-job-queue/internal/service"
	"time"
)

func main() {
	Q := model.QueueInit()
	reg := model.InitRegistry()
	limits := service.InitLimitFuncNumber()
	service.RegistryAllHandler(reg)
	Q.Push("print", []byte("hello"))
	go service.Worker(Q, reg, limits)
	time.Sleep(3 * time.Second)
	fmt.Println("end...")
}

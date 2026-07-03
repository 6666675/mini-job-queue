package service

import (
	"mini-job-queue/internal/handler"
	"mini-job-queue/internal/model"
)

func RegistryAllHandler(registry *model.Registry) {
	registry.Register(handler.PrintHandler{})
}

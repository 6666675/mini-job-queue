package handler

import (
	"fmt"
)

type PrintHandler struct{}

func (p PrintHandler) Name() string {
	return "print"
}
func (p PrintHandler) Run(Payload []byte) error {
	fmt.Println(string(Payload))
	return nil
}

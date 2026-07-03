package model

type Handler interface {
	Name() string
	Run(Payload []byte) error
}

package model

// 注册函数
type Registry struct {
	handlers map[string]Handler
}

func InitRegistry() *Registry {
	return &Registry{handlers: make(map[string]Handler)}
}
func (r *Registry) Register(h Handler) {
	r.handlers[h.Name()] = h
}
func (r *Registry) Get(name string) (Handler, bool) {
	h, ok := r.handlers[name] //ok代表是否存在键值，等同h!=nil判断
	return h, ok
}

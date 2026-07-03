package service

func InitLimitFuncNumber() map[string]chan struct{} {
	limits := map[string]chan struct{}{
		"print": make(chan struct{}, 2),
	}
	return limits
}

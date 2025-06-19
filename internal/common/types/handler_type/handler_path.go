package handler_type

type HandlerPath string

func (h HandlerPath) String() string {
	return string(h)
}

func (h HandlerPath) GetPath() string {
	return h.String()
}
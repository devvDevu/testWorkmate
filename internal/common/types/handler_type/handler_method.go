package handler_type

const (
	HandlerMethodGet HandlerMethod = "GET"
	HandlerMethodPost HandlerMethod = "POST"
	HandlerMethodPut HandlerMethod = "PUT"
	HandlerMethodDelete HandlerMethod = "DELETE"
)

type HandlerMethod string

func (h HandlerMethod) String() string {
	return string(h)
}

func (h HandlerMethod) GetMethod() string {
	return h.String()
}
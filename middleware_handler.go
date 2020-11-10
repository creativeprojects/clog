package clog

// middlewareHandler is a handler that sits in front of another handler
type middlewareHandler struct {
	next Handler
}

func newMiddlewareHandler(next Handler) *middlewareHandler {
	return &middlewareHandler{
		next: next,
	}
}

// SetHandler sets the next handler in the chain
func (h *middlewareHandler) SetHandler(handler Handler) {
	h.next = handler
}

// GetHandler returns the next handler in the chain
func (h *middlewareHandler) GetHandler() Handler {
	return h.next
}

package clog

import (
	"sync"
	"sync/atomic"
)

// overflowHandler is a MemoryHandler that calls overflow func when the size of
// recorded messages reaches or exceeds the specified overflowSize.
type overflowHandler struct {
	MemoryHandler
	overflow     func()
	overflowSize int64
}

func newOverflowHandler(overflow func(*overflowHandler), overflowSize int64) (handler *overflowHandler) {
	once := new(atomic.Pointer[sync.Once])
	once.Store(new(sync.Once))

	handleOverflow := func() {
		once.Load().Do(func() {
			defer once.Store(new(sync.Once))
			defer handler.Reset()
			overflow(handler)
		})
	}

	handler = &overflowHandler{
		MemoryHandler: *NewMemoryHandler(),
		overflow:      handleOverflow,
		overflowSize:  overflowSize,
	}
	return
}

// Transfers all buffered messages from the specified src to dest, if the src is an overflowHandler
func transferLogFromOverflowHandler(dst, src Handler) {
	if ofh, ok := src.(*overflowHandler); ok && src != nil {
		for ofh.TransferTo(dst) {
		}
	}
}

func (o *overflowHandler) LogEntry(entry LogEntry) error {
	defer func() {
		if o.Size() > o.overflowSize {
			o.overflow()
		}
	}()
	return o.MemoryHandler.LogEntry(entry)
}

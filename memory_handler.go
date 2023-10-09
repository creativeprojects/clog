package clog

import (
	"reflect"
	"sync"
	"sync/atomic"
)

// MemoryHandler save messages in memory (useful for unit testing).
type MemoryHandler struct {
	prefix string
	log    []string
	info   []memoryLogInfo
	size   atomic.Int64
	mu     sync.Mutex
}

type memoryLogInfo struct {
	start int
	level LogLevel
}

var memoryLogFixedSize = int64(reflect.TypeOf(memoryLogInfo{}).Size() + reflect.TypeOf("").Size())

// NewMemoryHandler creates a new MemoryHandler that keeps logs in memory.
func NewMemoryHandler() *MemoryHandler {
	return &MemoryHandler{
		log:  make([]string, 0, 10),
		info: make([]memoryLogInfo, 0, 10),
	}
}

// LogEntry keep the message in memory.
func (h *MemoryHandler) LogEntry(logEntry LogEntry) error {
	message := logEntry.GetMessage()
	{
		h.mu.Lock()
		defer h.mu.Unlock()

		if len(h.prefix) > 0 {
			message = h.prefix + message
		}
		h.size.Add(memoryLogFixedSize + int64(len(message)))
		h.log = append(h.log, message)
		h.info = append(h.info, memoryLogInfo{
			start: len(h.prefix),
			level: logEntry.Level,
		})
		return nil
	}
}

// SetPrefix adds a prefix to every log message.
// Please note no space is added between the prefix and the log message
func (h *MemoryHandler) SetPrefix(prefix string) Handler {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.prefix = prefix
	return h
}

// swap log between this and the other memory handler
func (h *MemoryHandler) swap(other *MemoryHandler) {
	h.mu.Lock()
	defer h.mu.Unlock()
	other.mu.Lock()
	defer other.mu.Unlock()

	h.log, other.log = other.log, h.log
	h.info, other.info = other.info, h.info
	other.size.Store(h.size.Swap(other.Size()))
	return
}

func (h *MemoryHandler) pop() (level LogLevel, startIdx int, message string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	tailIdx := len(h.log) - 1
	level = h.info[tailIdx].level
	startIdx = h.info[tailIdx].start
	message = h.log[tailIdx]

	h.size.Add(-(memoryLogFixedSize + int64(len(message))))
	h.log = h.log[:tailIdx]
	h.info = h.info[:tailIdx]
	return
}

// Pop returns the latest log from the internal storage (and removes it)
func (h *MemoryHandler) Pop() string {
	_, _, latest := h.pop()
	return latest
}

// PopWithLevel returns the latest log from the internal storage (and removes it)
func (h *MemoryHandler) PopWithLevel() (level LogLevel, prefix, message string) {
	var start int
	level, start, message = h.pop()
	prefix = message[:start]
	message = message[start:]
	return
}

// TransferTo transfers (and removes) all messages from the internal storage to another handler.
// Returns true as messages were transferred.
func (h *MemoryHandler) TransferTo(receiver Handler) bool {
	if receiver == h {
		return false
	}

	// swap log and release mutex before sending to receiver
	local := NewMemoryHandler()
	local.swap(h)

	// calldepth is not supported as the caller is TransferTo, skipping all stack frames
	const depth = 1_000

	if receiver != nil {
		for i, message := range local.log {
			info := local.info[i]
			if info.start > 0 {
				message = message[info.start:]
			}
			_ = receiver.LogEntry(NewLogEntry(depth, info.level, message))
		}
	}

	return len(local.log) > 0
}

// Reset clears internally buffered log
func (h *MemoryHandler) Reset() { h.swap(NewMemoryHandler()) }

// Logs return a list of all the messages sent to the logger
func (h *MemoryHandler) Logs() []string {
	h.mu.Lock()
	defer h.mu.Unlock()

	return append([]string(nil), h.log...)
}

// Empty return true when the internal list of logs is empty
func (h *MemoryHandler) Empty() bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	return len(h.log) == 0
}

// Size return the estimated amount of memory used by this handler
func (h *MemoryHandler) Size() int64 {
	return h.size.Load()
}

// Verify interface
var (
	_ Handler = &MemoryHandler{}
)

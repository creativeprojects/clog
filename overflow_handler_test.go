package clog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverflow(t *testing.T) {
	var logs []string
	handler := newOverflowHandler(func(handler *overflowHandler) {
		logs = handler.Logs()
	}, 1024)
	logger := NewLogger(handler)

	for retry := 0; retry < 3; retry++ {
		logs = nil
		line := 0
		for logs == nil {
			assert.Equal(t, handler.Size() == 0, handler.Empty())
			logger.Infof("-- %03d", line)
			line++
		}

		assert.Equal(t, 27, line)
		assert.True(t, handler.Empty())
		assert.Equal(t, int64(0), handler.Size())

		assert.NotNil(t, logs)
		for i := 0; i < line; i++ {
			assert.Equal(t, fmt.Sprintf("-- %03d", i), logs[i])
		}
	}
}

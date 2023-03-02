package clog

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogEntry(t *testing.T) {
	entry := NewLogEntry(2, LevelInfo, "a", "b", 10, true, false)
	assert.Equal(t, "ab10 true false", entry.GetMessage())
}

func TestNewLogEntryf(t *testing.T) {
	entry := NewLogEntryf(2, LevelInfo, "%s%s%d %v %v", "a", "b", 10, true, false)
	assert.Equal(t, "ab10 true false", entry.GetMessage())
}

func TestNewLogEntryFunc(t *testing.T) {
	entry := NewLogEntry(2, LevelInfo, fmt.Sprintf, "%s - %03d", "abc", 4)
	assert.Equal(t, "abc - 004", entry.GetMessage())
}

func TestMessageFromFuncCall(t *testing.T) {
	in := func(args ...interface{}) []interface{} { return args }
	xy := func(x, y int) string { return fmt.Sprintf("x = %d ; y = %d", x, y) }
	stringV := func(s ...string) string { return strings.Join(s, ", ") }
	stringA := func(s []string) string { return stringV(s...) }

	tests := []struct {
		in          []interface{}
		expect      string
		expectPanic bool
	}{
		{in: in(""), expect: "-"},
		{in: in(1), expect: "-"},
		{in: in(), expect: "-"},

		{in: in(xy), expectPanic: true},
		{in: in(xy, 2), expectPanic: true},
		{in: in(xy, 1, 2), expect: "x = 1 ; y = 2"},
		{in: in(xy, true, false), expectPanic: true},
		{in: in(xy, nil, nil), expectPanic: true},
		{in: in(xy, "y", "x"), expectPanic: true},
		{in: in(xy, 4, 1, 2), expect: "x = 4 ; y = 1"},

		{in: in(fmt.Sprintf), expectPanic: true},
		{in: in(fmt.Sprintf, "%02d:%02d"), expect: "%!d(MISSING):%!d(MISSING)"},
		{in: in(fmt.Sprintf, "%02d:%02d", 3, 4), expect: "03:04"},

		{in: in(stringV, 3), expectPanic: true},
		{in: in(stringV, "a", "b"), expectPanic: true},
		{in: in(stringV, []string{"a", "b"}), expect: "a, b"},
		{in: in(stringA, []string{"a", "b"}), expect: "a, b"},

		{in: in(func() string { return "fast path" }), expect: "fast path"},
		{in: in(func() (string, error) { return "abc", nil }), expect: "abc"},

		{in: in(func() error { return nil }), expect: "[<error Value>]"},
		{in: in(func() []int { return []int{0, 0} }), expect: "[<[]int Value>]"},
		{in: in(func() (int, bool) { return 5, false }), expect: "[<int Value> <bool Value>]"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if test.expectPanic {
				assert.Panics(t, func() { messageFromFuncCall(test.in) })
			} else {
				assert.Equal(t, test.expect, messageFromFuncCall(test.in))
			}
		})
	}
}

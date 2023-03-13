package clog

import (
	"fmt"
	"reflect"
)

// LogEntry represents a log entry
type LogEntry struct {
	Calldepth int           // Calldepth is used to calculate the right place where we called the log method
	Level     LogLevel      // Debug, Info, Warning or Error
	Format    string        // Format for *printf (leave blank for *print)
	Values    []interface{} // Values for *print and *printf OR a string func at index 0 with params from index 1
}

// NewLogEntry creates a new LogEntry composed of values.
//
// values parameter is comparable to fmt.Sprint(values...)
func NewLogEntry(callDepth int, level LogLevel, values ...interface{}) LogEntry {
	return LogEntry{
		Calldepth: callDepth,
		Level:     level,
		Values:    values,
	}
}

// NewLogEntryf creates a new formatted LogEntry with values.
//
// parameters are comparable to fmt.Sprintf(format, values...)
func NewLogEntryf(callDepth int, level LogLevel, format string, values ...interface{}) LogEntry {
	return LogEntry{
		Calldepth: callDepth,
		Level:     level,
		Format:    format,
		Values:    values,
	}
}

// GetMessage returns the formatted message from Format & Values
func (l LogEntry) GetMessage() string {
	if l.Format == "" {
		if vl := len(l.Values); vl > 0 {
			switch value := l.Values[0].(type) {
			case string:
				// fast path for string arg
				if vl == 1 {
					return value
				}
			case func() string:
				// fast path for simple string func
				return value()
			default:
				if reflect.TypeOf(value).Kind() == reflect.Func {
					return messageFromFuncCall(l.Values)
				}
			}
		}
		return fmt.Sprint(l.Values...)
	}
	return fmt.Sprintf(l.Format, l.Values...)
}

// GetMessageWithLevelPrefix returns the formatted message from Format & Values prefixed with the level name
func (l LogEntry) GetMessageWithLevelPrefix() string {
	return l.Level.String() + " " + l.GetMessage()
}

// messageFromFuncCall calls the func at funcAndArgs[0] using params from funcAndArgs[1:] and returns the result as string
func messageFromFuncCall(funcAndParams []interface{}) string {
	if len(funcAndParams) < 1 {
		return "-"
	}

	fn := reflect.ValueOf(funcAndParams[0])
	if fn.Kind() != reflect.Func {
		return "-"
	}

	fnType := fn.Type()
	funcAndParams = funcAndParams[1:]
	out, in := []reflect.Value(nil), make([]reflect.Value, 0, fnType.NumIn())

	if fnType.IsVariadic() {
		// non-Variadic params
		paramCount := fnType.NumIn() - 1
		for i := 0; i < paramCount; i++ {
			in = append(in, reflect.ValueOf(funcAndParams[i]))
		}
		// Variadic part
		if fnType.In(paramCount).Elem().Kind() == reflect.Interface {
			in = append(in, reflect.ValueOf(funcAndParams[paramCount:])) // func's variadic is "...any" > use remaining params directly
		} else {
			in = append(in, reflect.ValueOf(funcAndParams[paramCount])) // func's variadic is another type > correct array must be passed as last arg
		}
		out = fn.CallSlice(in)
	} else {
		// func params
		for _, param := range funcAndParams[:fnType.NumIn()] {
			in = append(in, reflect.ValueOf(param))
		}
		out = fn.Call(in)
	}

	if len(out) >= 1 && out[0].Kind() == reflect.String {
		return out[0].String()
	} else {
		return fmt.Sprint(out)
	}
}

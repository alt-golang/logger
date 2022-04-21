package logger

import (
	"testing"
)

func TestLogging(t *testing.T) {

	logger := GetFactory().GetLogger("github.com/alt-golang/logger")
	logger.Info("Hello")
	logger.Trace("Trace")
	logger.Error("Error")
	meta := map[string]string{
		"map": "hello",
	}
	logger.ErrorWithMetadata("Error", meta)

	d := struct {
		SomeData  string
		otherData string
	}{
		SomeData:  "This is it",
		otherData: "sdfsdf",
	}
	logger.ErrorWithMetadata("Error", &d)
	logger.ErrorWithMetadata("Error", d)
	trace := logger.IsTraceEnabled()

	t.Log(trace)
	if logger.IsTraceEnabled() != false {
		t.Errorf("logger.IsTraceEnabled() == true: logger.IsTraceEnabled(): %t", false)
	}
}

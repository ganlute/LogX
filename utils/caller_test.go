package utils

import (
	"strings"
	"testing"
)

func TestGetCallerName(t *testing.T) {
	name := GetCallerName(0)
	if strings.Contains(name, "caller_test.go:9") {
		// pass
	} else {
		t.Error("name != caller_test.go:9")
	}
}

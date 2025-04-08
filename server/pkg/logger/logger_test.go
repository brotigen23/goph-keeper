package logger

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New().Testing()
	err := fmt.Errorf("Error")
	logger.Error(err)
}

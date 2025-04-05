package logger

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New().Default()
	err := fmt.Errorf("AAAAAAAAa")
	logger.Error(err)
}

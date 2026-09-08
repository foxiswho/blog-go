package test

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := log2.NewDefault(log2.LevelDebug)
	logger.Infof("hello %s", "world")
}

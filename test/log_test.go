package test

import (
	"github.com/foxiswho/blog-go/pkg/log2"
	"testing"
)

func TestLog(t *testing.T) {
	logger := log2.NewDefault(log2.LevelDebug)
	logger.Infof("hello %s", "world")
}

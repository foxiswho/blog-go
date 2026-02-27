package test

import (
	"fmt"
	"github.com/pangu-2/go-tools/tools/datetimePg"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ti := time.Now()
	tmp := datetimePg.Format(ti, "2006")
	fmt.Printf("%s\n", tmp)
}

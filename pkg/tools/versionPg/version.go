package versionPg

import (
	"github.com/pangu-2/go-tools/tools/datetimePg"
	"strings"
	"time"
)

func Make() string {
	now := time.Now()
	format := now.Format(datetimePg.YMDHIS_SSS)
	format = strings.Replace(format, ".", "", -1)
	return format
}

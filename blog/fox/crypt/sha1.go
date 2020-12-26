package crypt

import (
	"io"
	"fmt"
	"crypto/sha1"
)

func Sha1(msg string) string {
	h := sha1.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

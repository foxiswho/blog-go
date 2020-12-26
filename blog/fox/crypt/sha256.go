package crypt

import (
	"io"
	"fmt"
	"crypto/sha256"
)

func Sha256(msg string) string {
	h := sha256.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}

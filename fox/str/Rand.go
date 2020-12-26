package str

import (
	"math/rand"
	"time"
)
// randseed
func GetRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
// rand salt
func RandSalt() string {
	var salt = ""
	for i := 0; i < 4; i++ {
		ran := GetRand()
		salt += string(SALT[ran.Intn(len(SALT))])
	}
	return salt
}
const (
	SALT = "$^*#,.><)(_+f*m"
)

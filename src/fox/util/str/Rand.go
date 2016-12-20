package str

import (
	"time"
	"math/rand"
)
// randseed
func GetRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
// rand salt
func RandSalt() string {
	var salt = ""
	for i := 0; i < 4; i++ {
		rand := GetRand()
		salt += string(SALT[rand.Intn(len(SALT))])
	}
	return salt
}
const (
	SALT = "$^*#,.><)(_+f*m"
)
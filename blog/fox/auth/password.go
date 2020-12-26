package auth

import (
	"github.com/foxiswho/blog-go/blog/fox/config"
	"github.com/foxiswho/blog-go/blog/fox/crypt"
)

// encrypt password
func PasswordSalt(pass, salt string) string {
	salt1 := "4%$@w"
	password_salt := config.String("password_salt")
	str :=salt1+pass+salt+password_salt
	//return crypt.Md5(crypt.Sha256(str))
	return crypt.Md5(str)
}
//验证
func PasswordVerify(password,  pass, salt string) bool {
	return password == PasswordSalt(pass, salt)
}

package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type Aes struct {
}

func (t *Aes) getKey() []byte {
	strKey := "1234567890123456"
	keyLen := len(strKey)
	if keyLen < 16 {
		fmt.Println("res key 长度不能小于16")
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func (t *Aes) Encrypt(strMesg string) ([]byte, error) {
	key := t.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}

//解密字符串
func (t *Aes) Decrypt(src []byte) (str1 string, err  error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
			//return "",err
		}
	}()
	key := t.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}
package utils

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func ScryptPW(password string) (string, error) {
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		fmt.Printf("cryptPW fail%s", err)
		return " ", err
	}
	psd := base64.StdEncoding.EncodeToString(dk)
	return string(psd), nil
}

package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Encrypt 密码加密
func Encrypt(source string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashPwd), err
}

// Compare 密码比对 (传入未加密的密码即可)
func Compare(hashedPassword, password string) error {
	if hashedPassword == password {
		return nil
	}
	return fmt.Errorf("password is invalid")
	// return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

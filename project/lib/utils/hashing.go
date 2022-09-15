package utils

import "golang.org/x/crypto/bcrypt"

type Hashing struct {
}

func NewHashing() *Hashing {
	return &Hashing{}
}

func (h Hashing) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h Hashing) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

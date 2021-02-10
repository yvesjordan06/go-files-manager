package utilities

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	bytePwd := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)

	if err != nil {
		log.Fatalln("Security Error: Couldn't hash the password")
	}

	return string(hash)
}

func CompareHashPassword(plain, hashed string) bool {
	plainByte := []byte(plain)
	hashedByte := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(hashedByte, plainByte)

	if err != nil {
		return false
	}
	return true
}

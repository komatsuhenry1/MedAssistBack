package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"unicode"
	"golang.org/x/crypto/bcrypt"
)

const passwordLength = 10

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"!@#$%&*_-+="

func GeneratePassword() (string, error) {
	fmt.Println("entrou na gemereate pawssword")
	password := make([]byte, passwordLength)

	for i := range password {
		char, err := randomChar(charset)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	fmt.Println("password", password)

	return string(password), nil
}

func randomChar(charset string) (byte, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, err
	}
	return charset[index.Int64()], nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func ValidatePassword(password string) bool {
	hasSpecial := false
	hasNumber := false
	hasLetter := false

	for _, char := range password {
		fmt.Println(string(char))

		if unicode.IsDigit(char) {
			hasNumber = true
		}
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsPunct(char) {
			hasSpecial = true
		}
	}

	return hasSpecial && hasNumber && hasLetter
}

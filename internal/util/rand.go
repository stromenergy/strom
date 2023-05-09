package util

import (
	"math/rand"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func RandomString(length int) string {
	str := make([]rune, length)

	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}

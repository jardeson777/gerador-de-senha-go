package utils

import "math/rand"

func GenerateLetterUppercase() string {
	return string(rune(65 + rand.Intn(26)))
}

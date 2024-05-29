package utils

import "math/rand"

func GenerateLetterLowcase() string {
	return string(rune(97 + rand.Intn(26)))
}

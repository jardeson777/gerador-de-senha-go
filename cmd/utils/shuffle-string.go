package utils

import "math/rand"

func ShuffleString(s string) string {
	r := []rune(s)
	for i := range r {
		j := rand.Intn(i + 1)
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
package utils

import (
	"math/rand"
	"strconv"
)

func GenerateNumber() string {
	return strconv.Itoa(rand.Intn(10))
}
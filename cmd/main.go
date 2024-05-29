package main

import (
	"fmt"
	"gerador-de-senha-go/cmd/password"
)

func main() {
    fmt.Println(password.Generate(9, 3, 3, 3))
}

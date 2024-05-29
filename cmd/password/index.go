package password

import "gerador-de-senha-go/cmd/utils"

func Generate(lenght int, amountNumber int, amountLetterLowcase int, amountLetterUppercase int) string {
	var password string

	if lenght == 0 || lenght < 0 || lenght != (amountNumber + amountLetterLowcase + amountLetterUppercase) {
		lenght = 8
	}

	for i := 0; i < amountNumber; i++ {
		password += utils.GenerateNumber()
	}

	for i := 0; i < amountLetterUppercase; i++ {
		password += utils.GenerateLetterUppercase()
	}

	for i := 0; i < amountLetterLowcase; i++ {
		password += utils.GenerateLetterLowcase()
	}

	password = utils.ShuffleString(password)

	return password
}

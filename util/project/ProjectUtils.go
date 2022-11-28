package project

import (
	"fmt"
	"math/rand"
)

func GenerateProjectID(projectName string, generateRandomID bool) string {
	if generateRandomID {
		lowerCase := "abcdefghijklmnopqrstuvwxyz"
		numbers := "0123456789"

		letters := []rune(lowerCase + numbers)
		b := make([]rune, 4)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}

		return fmt.Sprintf("%s-%s", projectName, string(b))
	}

	return projectName
}

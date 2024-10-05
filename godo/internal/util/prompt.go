package util

import (
	"bufio"
	"fmt"
	"os"
)

// Prompts the user if they would like to proceed by asking Y/n
func PromptUser(prompt *string) (userResponse bool, err error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", *prompt)
	fmt.Print("Would you like to proceed? [Y/n] ")

	char, _, err := reader.ReadRune()
	if err != nil { // Rune could not be encoded to UTF-8
		return false, fmt.Errorf("rune could not be encoded to UTF-8")
	}

	switch char {
	case 'Y':
	case 'y':
		return true, nil

	case 'N':
	case 'n':
		return false, nil
	}

	return false, err
}

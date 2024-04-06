package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"unicode"
)

// assuming input comes from Text file with just credit card numbers

func hasFourConsecutiveDigits(cardNumber string) bool {
	count := 1

	for i := 1; i < len(cardNumber); i++ {
		if unicode.IsDigit(rune(cardNumber[i])) && cardNumber[i] == cardNumber[i-1] {
			count++
			if count >= 4 {
				return true
			}
		} else {
			count = 1
		}
	}

	return false
}

func removeDashes(cardNumber string) string {
	chars := []rune(cardNumber)

	for i := 4; i < 16; i += 4 {
		if chars[i] == '-' {
			chars = append(chars[:i], chars[i+1:]...)
		} else {
			return string(chars)
		}
	}
	return string(chars)
}

func isValid(cardNumber string) bool {
	regex16 := `^\d{16}$`
	start456 := `^[456]`
	only_16_digits, _ := regexp.MatchString(regex16, cardNumber)
	starts_with_456, _ := regexp.MatchString(start456, cardNumber)

	return only_16_digits && starts_with_456

}

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		println(line)
		line = removeDashes(line)
		if isValid(line) && !hasFourConsecutiveDigits(line) {
			println("Valid")
		} else {
			println("Invalid")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}

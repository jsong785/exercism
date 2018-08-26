package luhn

import (
	"strconv"
	"unicode"
)

func Valid(s string) bool {

	sum := 0

	numFound := 0
	altNumberFlag := false
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])

		if unicode.IsSpace(c) {
			continue
		}

		if !unicode.IsDigit(c) {
			return false
		}

		number, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}

		if altNumberFlag {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}

		altNumberFlag = !altNumberFlag
		sum += number
		numFound++
	}

	return numFound > 1 && (sum%10) == 0
}

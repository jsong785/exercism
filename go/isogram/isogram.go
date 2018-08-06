package isogram

import "unicode"

func HasChar(list[]rune, char rune) bool {
    for _, c := range list {
        if c == char {
            return true
        }
    }
    return false
}

func IsIsogram(input string) bool {
    var charList []rune

    for _, c := range input {
        if !unicode.IsLetter(c) {
            continue
        }

        c = unicode.ToUpper(c)
        if HasChar(charList, c) {
            return false
        }
        charList = append(charList, c)
    }

    return true
}

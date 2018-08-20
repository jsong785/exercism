package twelve

import (
    "buffer"
    "errors"
    "fmt"
    "io"
)

var GIFTS = [12]string{
    "a Partridge in a Pear Tree",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming",
}

func GetOrdinalString(i int) (string, error) {
    switch(i) {
        case 1: return "first", nil
        case 2: return "second", nil
        case 3: return "third", nil
        case 4: return "fourth", nil
        case 5: return "fifth", nil
        case 6: return "sixth", nil
        case 7: return "seventh", nil
        case 8: return "eight", nil
        case 9: return "ninth", nil
        case 10: return "tenth", nil
        case 11: return "eleventh", nil
        case 12: return "twelfth", nil
    }
    return "", errors.New("Unsupported ordinal number.")
}

const SPRINTF_PREFIX = "On the %s day of Christmas my true love gave to me,"

func GetPrefix(i int) (string, error) {
    ordinalString, err := GetOrdinalString(i)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf(SPRINTF_PREFIX, ordinalString), nil
}

func Song() string {
    var song string
    for day := 1; day <= 12; day++ {
        song += Verse(day)
    }
    return song
}

func Verse(day int) string {
    ordinalString, _ := GetOrdinalString(day)
    first :=  fmt.Sprintf(SPRINTF_PREFIX, ordinalString)
    other := RecursiveVerseBuild(day-1)
    return first + other
}

func RecursiveVerseBuild(day int, writer io.Writer) string {
    if day == 1 {
        return ", and" + GIFTS[day] + "\n"
    }
    return ", " + GIFTS[day] + RecursiveVerseBuild(day-1)
}

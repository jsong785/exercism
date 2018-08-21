package twelve

import (
    "bytes"
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
        case 8: return "eighth", nil
        case 9: return "ninth", nil
        case 10: return "tenth", nil
        case 11: return "eleventh", nil
        case 12: return "twelfth", nil
    }
    return "", errors.New("Unsupported ordinal number.")
}

const SPRINTF_PREFIX = "On the %s day of Christmas my true love gave to me, %s"
func GetVerseStart(day int) (string, error) {
    ordinalString, err := GetOrdinalString(day)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf(SPRINTF_PREFIX, ordinalString, GIFTS[day-1]), nil
}

func Song() string {
    buffer := bytes.Buffer{}

    for day := 1; day <= 12; day++ {
        song := Verse(day)
        buffer.WriteString(song)
        buffer.WriteString("\n")
    }

    return buffer.String()
}

func Verse(day int) string {
    buffer := bytes.Buffer{}

    firstVerse, _ := GetVerseStart(day)
    buffer.WriteString(firstVerse)

    if needsMoreGifts := (day-1 > 0); needsMoreGifts {
        RecursiveVerseBuild(day-1, &buffer)
    }

    buffer.WriteString(".")
    return buffer.String()
}

func RecursiveVerseBuild(day int, writer io.Writer) {
    if day == 1 {
        verse := ", and " + GIFTS[day-1]
        writer.Write([]byte(verse))
        return
    }
    verse := ", " + GIFTS[day-1]
    writer.Write([]byte(verse))
    RecursiveVerseBuild(day-1, writer)
}

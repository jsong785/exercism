package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(num int) string {
	isPling := (num % 3) == 0
	isPlang := (num % 5) == 0
	isPlong := (num % 7) == 0

	if !isPling &&
		!isPlang &&
		!isPlong {
		return strconv.Itoa(num)
	}

	var buffer bytes.Buffer
	if isPling {
		buffer.WriteString("Pling")
	}
	if isPlang {
		buffer.WriteString("Plang")
	}
	if isPlong {
		buffer.WriteString("Plong")
	}
	return buffer.String()
}

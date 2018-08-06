package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return -1, errors.New("strings are not equal length.")
    }
    
    var distance int = 0
    for i := 0; i < len(a); i++ {
        charA := a[i]
        charB := b[i]

        if charA != charB {
            distance++
        }
    }
    return distance, nil
}

package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
    m := FreqMap{}
    for _, r := range s {
        m[r]++
    }
    return m
}

func ConcurrentFrequency(list []string) FreqMap {
    ch := make(chan FreqMap)
    ConcurrentFrequencyStart(ch, list)
    return ConcurrentFrequencyGatherResults(ch, uint(len(list)))
}

func ConcurrentFrequencyStart(ch chan FreqMap, list[] string) {
    for i := 0; i < len(list); i++ {
        go func (s string) {
            resultFreqMap := Frequency(s)
            ch <- resultFreqMap
        }(list[i])
    }
}

func ConcurrentFrequencyGatherResults(ch chan FreqMap, expectedCount uint) FreqMap {
    m := FreqMap{}

    for i := uint(0); i < expectedCount; i++ {
        for char, charCount := range <-ch {
            m[char] += charCount
        }
    }

    return m
}

package diffsquares

func SumOfFirstNPositiveIntegers(n int) int {
    // n(n+1)
    // -----
    //   2
    numerator := n * (n + 1)
    return numerator / 2
}

func SumOfSquaresFirstNPositiveIntegers(n int) int {
    // n(n + 1)(2n + 1)
    // ---------------
    //        6
    numerator := n * (n + 1) * ((2 * n) + 1)
    return numerator / 6
}

func SquareOfSums(n int) int {
    sum := SumOfFirstNPositiveIntegers(n)
    return sum * sum
}

func SumOfSquares(n int) int {
    return SumOfSquaresFirstNPositiveIntegers(n)
}

func Difference(n int) int {
    return SquareOfSums(n) - SumOfSquares(n)
}

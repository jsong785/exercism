package grains

import "errors"

func TwoPowerN(n uint64) uint64 {
	return 1 << n
}

func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return 0, errors.New("Invalid square number.")
	}
	exp := uint64(square - 1)
	return TwoPowerN(exp), nil
}

func Total() uint64 {
	// speed up
	// 2^0 + 2^1 + ... 2^n = 2^(n+1) - 1

	// 0 -> (63 + 1)
	return (1 << 64) - 1
}

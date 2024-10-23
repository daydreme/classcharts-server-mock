package util

import (
	"crypto/rand"
	"math/big"
)

const IdMaxDigits = 6

func RandomInt(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}

func RandomId() int {
	id, err := RandomInt(10^IdMaxDigits, 9*(10^IdMaxDigits))
	if err != nil {
		panic(err)
	}

	return id
}

package rand

import (
	"math"
	"math/rand"
)

type RandomInt int

func NewRandomInt(n int) RandomInt {
	rand.Seed(42)
	return RandomInt(rand.Intn(n))
}

func (r RandomInt) Int() int {
	return int(r)
}

func (r RandomInt) Float() float64 {
	return float64(r)
}

func (r RandomInt) sqrt() int {
	return int(math.Sqrt(r.Float()))
}

func (r RandomInt) IsPrime() bool {
	if r == 0 || r == 1 {
		return false
	}
	for i := 2; i < r.sqrt(); i++ {
		if r.Int()%i == 0 {
			return false
		}
	}
	return true
}

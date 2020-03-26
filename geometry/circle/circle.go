package circle

import (
	"go_learn/calculate"
	"math"
)

func Area(r float32) float32 {
	return math.Pi * calculate.Multiply(r, r)
}

func Perimeter(r float32) float32 {
	return 2 * calculate.Multiply(2, calculate.Multiply(math.Pi, r))
}

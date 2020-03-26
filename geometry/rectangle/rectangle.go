package rectangle

import "go_learn/calculate"

func Area(a, b float32) float32 {
	return calculate.Multiply(a, b)
}

func Perimeter(a, b float32) float32 {
	return calculate.Multiply(calculate.Add(a, b), 2)
}

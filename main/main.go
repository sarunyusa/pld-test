package main

import (
	"math"
)

//โจทย์: เช็คว่า integer เป็น palindrome หรือเปล่า โดยห้ามแปลง integer เป็น string
//ตัวอย่าง
//
//Integer = 121
//Expected return = true
//
//Integer = -121
//Expected return = false // -121 != 121-
//
//Integer = 10
//Expected return = false // 10 != 01

func main() {

}

func Palindrome(i int) bool {
	if i < 0 {
		return false
	}

	return i == Reverse(i)
}

func Reverse(in int) int {
	is := make([]int,0)
	mn := in
	i := 1
	sum := 0.0

	var f func(f float64) float64
	if in > 0 {
		f = math.Floor
	} else {
		f = math.Ceil
	}

	for mn != 0 {
		mn = int(f(float64(in) / math.Pow10(i)))

		m := math.Mod(float64(in), math.Pow10(i))
		m = f(m / math.Pow10(i-1))
		is = append(is, int(m))
		sum += m
		i++
	}

	l := len(is)
	result := 0
	for i = range is {
		result += is[l-i-1] * int(math.Pow10(i))
	}

	return result
}

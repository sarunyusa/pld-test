package main

import (
	"fmt"
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

func Palindome(i int) bool {
	if i < 0 {
		return false
	}

	return i == Reverse(i)
}

func Reverse(in int) int {
	is := make([]int,0)
	mn := -1
	i := 1
	for mn != 0 {
		fmt.Println("in", in)
		fmt.Println("mn",mn)
		fmt.Println("i",i)
		fmt.Println(len(is))
		fmt.Println(is)
		//fmt.Println(float64(in))
		//fmt.Println(float64(10^i))
		//fmt.Println(10^i)
		//fmt.Println(float64(in)/math.Pow(float64(10)/float64(i)))
		fmt.Println("----")
		mn = int(math.Floor(float64(in)/math.Pow10(i)))
		m := math.Mod(float64(in), math.Pow10(i))
		if i != 1 {
			fmt.Println(is)
			fmt.Println(is[i-1])
			m = m - float64(is[i-1])
		}
		is = append(is, int(m))
		fmt.Println("mn",mn)
		fmt.Println("m",m)
		i++
	}

	fmt.Println("is",is)

	l := len(is)
	fmt.Println("l",l)
	result := 0
	for i = range is {
		result += is[l-i] * (10^i)
	}

	return result
}

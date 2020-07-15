package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"strings"
)

//Цикл FOR

func Sqrt(x float64) float64 {
	z0 := 0.0

	for i := 0; i < 2; i++ {
		z0 = x

		x = z0 - (z0 * z0 - x)/ ( 2 * z0 )
		fmt.Println(x)
	}
	return x
}

//Срезы

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8,dy)

	for i := range slice {
		slice[i] = make([]uint8, dx)
	}

	for i := range slice {
		for j := range slice[i] {
			slice[i][j] = uint8(i ^ j)
		}
	}

	return slice
}

//maps

func WordCount(s string) map[string]int {

	str := strings.Fields(s)

	resWords := make(map[string]int)

	for i := range str{
		resWords[str[i]]++
	}

	return resWords
}

func fibonacci() func() int {

	first  := -1
	second := 1

	return func() int {
		res    := first + second
		first  = second
		second = res
		return res
	}
}

func main() {
	fmt.Println(Sqrt(9))

	pic.Show(Pic)

	wc.Test(WordCount)

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
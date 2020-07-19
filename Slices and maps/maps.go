package main

import (
	"fmt"
	"strings"
)

//кол-во слов в строке
func countWords(s string) map[string]int{
	count := make(map[string]int)
	for _, value := range strings.Fields(s) {
		count[value]++
	}

	return count
}
//2.Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
func numsInMap(v []int) map[int]int{

	mape := make(map[int]int)
	for _, value := range v{
		mape[value]++
	}

	return mape
}
//3.Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих
func checkDuplicates(v []int, b []int) []int{
	mape := make(map[int]bool)

	mass := make([]int, 0, 100)

	for _, value := range v {
		mape[value] = true
	}

	for _, value := range b{
		if _, ok := mape[value]; ok {
			mass = append(mass, value)
		}
	}

	return mass
}
//4.Сделать Фибоначчи с мемоизацией
func memes(memeMap map[int]int,i int) int{
	switch {
	case i == 0:
		return 0
	case i == 0 || i <= 2:
		return 1
	case memeMap[i] == 0:
		memeMap[i] = memes(memeMap, i-1) + memes(memeMap, i-2)
	}

	return memeMap[i]
}

func main() {

	memeMap := make(map[int]int)

	mapString := "cat fat rat dog cat cat house house"

	sliceForMap := []int{1,2,3,4,1,4,3,3,6,5,4,7,8,9,0,5,88,6,5}
	//1. Найти кол-во повторений слов в строке
	fmt.Println(countWords(mapString))
	//2.Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
	numsinmap := numsInMap(sliceForMap)
	fmt.Println(numsinmap)
	//3.Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих
	mass := checkDuplicates([]int{1,2,3,4,5,6,7,8,9}, []int{0,5,7})
	fmt.Println(mass)
	//4.Сделать Фибоначчи с мемоизацией
	for i := 0; i<5; i++  {
		fibonacci := memes(memeMap, i)
		fmt.Println(fibonacci)
	}
}

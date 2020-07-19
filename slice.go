package main

import (
	"fmt"
	"sort"
	"strings"
)
//К каждому элементу []int прибавить 1
func sliceAdd1(v []int) []int {

	for i:= 0; i < len(v); i++ {
		v[i]++
	}
	return v
}

//Добавить в конц число
func sliceAddEnd(v []int, num int) []int{
	v = append(v, num)

	return  v
}

//Добавить в начало число
func sliceAddBeg(v []int, num int) []int{
	tempSlice := []int{num}
	v = append(tempSlice, v...)

	return v

}

//Удалить последний элемент и вернуть его пользователю
func removeEl(v []int) ([]int, int){

	var i int

	i = v[len(v)-1]

	return v[:len(v)-1],i
}

func removeElBeg(v []int)  ([]int, int){
	var i int

	i = v[0]

	return v[1:], i
}
//Взять i-е число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передает пользователь
func removeIndexSlice(v []int, index int) ([]int, int){
	var i int

	i = v[index]
	v = append(v[:index], v[index+1:]...)

	return v, i
}
//Объединить два slice и вернуть новый со всеми элементами первого и второго
func unificationSlices(v []int, b []int) []int{

	return append(v, b...)
}
//Удалить числа которые есть во втором массиве
func removeDuplicates(v []int, b []int) []int{

	for i := len(v)-1; i>=0; i-- {
		for j := len(b)-1; j>=0; j--{
			if b[j] == v[i] {
				v = append(v[:i], v[i+1:]...)
				i--
			}
		}
	}

	return v
}
//сдвиг влево на 1
func leftSwap(v []int) []int{
	v = append(v[1:], v[:1]...)

	return v
}
//сдвиг влево на n
func leftSwapRandomVal(v []int, num int) []int{
	v = append(v[num:], v[:num]...)

	return v
}
//сдвиг вправо на 1
func rightSwap(v []int) []int{
	v = append(v[len(v)-1:], v[:len(v)-1]...)

	return v
}
//сдвиг вправо на n
func rightSwapRandomVal(v []int, num int) []int{
	v = append(v[len(v)-num:], v[:len(v)-num]...)

	return v
}
//вовзращаю копию слайса
func copySlice(v []int) []int{
	tempSlice :=make([]int,0,cap(v))
	copy(tempSlice,v)

	return v
}

func parity(v []int) []int{

	for index, _ := range v {
		if index%2 == 0  && index != len(v)-1{
			v[index], v[index+1] = v[index+1], v[index]
		}
	}

	return  v
}

func sortASlice(v []int) ([]int) {
	sort.Ints(v)

	return v
}

func sortDSlice(v []int) []int {
	sort.Ints(v)
	tempSlice := make([]int, 0, cap(v))

	for i := len(v)-1; i >= 0; i-- {
		tempSlice = append(tempSlice, v[i])
	}

	return tempSlice
}

func sortStrslice (s []string) []string{
	sort.Strings(s)

	return s
}
//кол-во слов в строке
func countWords(s string) map[string]int{
	count := make(map[string]int)
	for _, value := range strings.Fields(s) {
		count[value]++
	}

	return count
}

func numsInMap(v []int) map[int]int{

	mape := make(map[int]int)
	for _, value := range v{
		mape[value]++
	}

	return mape
}

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

	var i int
	
	memeMap := make(map[int]int)

	intSlice := []int{1,2,3,4,5,6,7,8,9}

	stringSlice := []string{"cat", "rat", "fat", "dead", "cut"}

	sliceForRemove := []int{1,2,3,4,5,6,7,8,9}

	mapString := "cat fat rat dog cat cat house house"

	sliceForMap := []int{1,2,3,4,1,4,3,3,6,5,4,7,8,9,0,5,88,6,5}

	v := sliceAdd1(intSlice)
	fmt.Println(v)

	end := sliceAddEnd(intSlice, 5)
	fmt.Println(end)

	beg := sliceAddBeg(intSlice, 5)
	fmt.Println(beg)

	sliceForRemove, i = removeEl(sliceForRemove)
	fmt.Println(sliceForRemove, i)

	sliceForRemove, i = removeElBeg(sliceForRemove)
	fmt.Println(sliceForRemove,i)

	sliceForRemove, i = removeIndexSlice(sliceForRemove, 3)
	fmt.Println(sliceForRemove, i)

	unionSlice := unificationSlices(intSlice, sliceForRemove)
	fmt.Println(unionSlice)

	withoutDuplicates := removeDuplicates([]int{1,2,3,4,7,10}, []int{2,4,10})
	fmt.Println(withoutDuplicates)

	left := leftSwap([]int{1,2,3,4,5})
	fmt.Println(left)

	leftRand := leftSwapRandomVal([]int{1,2,3,4,5,6,7}, 4)
	fmt.Println(leftRand)

	right := rightSwap([]int{1,2,3,4,5,6})
	fmt.Println(right)

	rightRand := rightSwapRandomVal([]int{1,2,3,4,5,6,7,8}, 4)
	fmt.Println(rightRand)

	fmt.Println(copySlice(intSlice))

	fmt.Println(parity(intSlice))

	ascendigSlice := sortASlice(intSlice)
	fmt.Println(ascendigSlice)

	descendingSlice := sortDSlice(intSlice)
	fmt.Println(descendingSlice)

	stringSort := sortStrslice(stringSlice)
	fmt.Println(stringSort)

	fmt.Println(countWords(mapString))

	numsinmap := numsInMap(sliceForMap)
	fmt.Println(numsinmap)

	mass := checkDuplicates([]int{1,2,3,4,5,6,7,8,9}, []int{0,5,7})
	fmt.Println(mass)

	for i := 0; i<5; i++  {
		fibonacci := memes(memeMap, i)
		fmt.Println(fibonacci)
	}

}

package main

import (
	"fmt"
	"sort"
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


func main() {

	var i int

	intSlice := []int{1,2,3,4,5,6,7,8,9}

	stringSlice := []string{"cat", "rat", "fat", "dead", "cut"}

	sliceForRemove := []int{1,2,3,4,5,6,7,8,9}

	//1 К каждому элементу []int прибавить 1
	v := sliceAdd1(intSlice)
	fmt.Println(v)
	//2 Добавить в конец slice число 5
	end := sliceAddEnd(intSlice, 5)
	fmt.Println(end)
	//3 Добавить в начало slice число 5
	beg := sliceAddBeg(intSlice, 5)
	fmt.Println(beg)
	//4 Взять последнее число slice, вернуть его пользователю, а из slice этот элемент удалить
	sliceForRemove, i = removeEl(sliceForRemove)
	fmt.Println(sliceForRemove, i)
	//5 Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить
	sliceForRemove, i = removeElBeg(sliceForRemove)
	fmt.Println(sliceForRemove,i)
	//6 Взять i-е число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передает пользователь
	sliceForRemove, i = removeIndexSlice(sliceForRemove, 3)
	fmt.Println(sliceForRemove, i)
	//7 Объединить два slice и вернуть новый со всеми элементами первого и второго
	unionSlice := unificationSlices(intSlice, sliceForRemove)
	fmt.Println(unionSlice)
	//8 Из первого slice удалить все числа, которые есть во втором
	withoutDuplicates := removeDuplicates([]int{1,2,3,4,7,10}, []int{2,4,10})
	fmt.Println(withoutDuplicates)
	//9 Сдвинуть все элементы slice на 1 влево.
	left := leftSwap([]int{1,2,3,4,5})
	fmt.Println(left)
	//10 Тоже, но сдвиг на заданное пользователем i
	leftRand := leftSwapRandomVal([]int{1,2,3,4,5,6,7}, 4)
	fmt.Println(leftRand)
	//11 Тоже, что 9, но сдвиг вправо
	right := rightSwap([]int{1,2,3,4,5,6})
	fmt.Println(right)
	//12 Тоже, но сдвиг на i
	rightRand := rightSwapRandomVal([]int{1,2,3,4,5,6,7,8}, 4)
	fmt.Println(rightRand)
	//13 Вернуть пользователю копию переданного slice
	fmt.Println(copySlice(intSlice))
	//14 В slice поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
	fmt.Println(parity(intSlice))
	//15 прямое упорядочивание
	ascendigSlice := sortASlice(intSlice)
	fmt.Println(ascendigSlice)
	//Обратное упорядовчивание
	descendingSlice := sortDSlice(intSlice)
	fmt.Println(descendingSlice)
	//Лексикографическое упорядочивание
	stringSort := sortStrslice(stringSlice)
	fmt.Println(stringSort)



}

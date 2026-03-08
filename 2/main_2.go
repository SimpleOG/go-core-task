package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(arr []int) []int {
	res := make([]int, 0, len(arr))
	for _, v := range arr {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func addElements(arr []int, num int) []int {
	res := arr
	res = append(res, num)
	return res
}
func CopySlice(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}
func removeElement(arr []int, index int) []int {
	if index < 0 {
		index = 0
	}
	if index > len(arr) {
		index = len(arr) - 1
	}
	res := make([]int, 0, cap(arr)-1)
	res = append(res, arr[:index]...)
	res = append(res, arr[index+1:]...)
	return res
}
func main() {
	//1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений,
	//которые генерируются случайным
	//образом (при каждом запуске должны получаться новые значения)
	originalSlice := make([]int, 10, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}
	fmt.Println(originalSlice)

	//2. Напишите функцию sliceExample, которая принимает
	//слайс и возвращает новый слайс,
	//содержащий только четные числа из исходного слайса.
	fmt.Println(sliceExample(originalSlice))

	//3. Напишите функцию addElements, которая принимает слайс и число.
	//Функция должна добавлять это число в конец слайса и возвращать новый слайс.
	num := 10
	newSlice := addElements(originalSlice, num)
	fmt.Printf("original slice: %v , len=%v, new slice: %v, len=%v \n",
		originalSlice, len(originalSlice), newSlice, len(newSlice))

	//4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию.
	// Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
	copiedSlice := CopySlice(originalSlice)
	copiedSlice[0] = 1
	originalSlice[3] = rand.Intn(100)
	fmt.Printf("original slice: %v, copied slice: %v \n", originalSlice, copiedSlice)

	//5. Напишите функцию removeElement, которая принимает слайс и индекс элемента,
	//который нужно удалить. Функция должна возвращать новый слайс
	//без элемента по указанному индексу.
	croppedSlice := removeElement(originalSlice, 5)
	fmt.Printf("original slice: %v, cropped slice: %v \n", originalSlice, croppedSlice)
	croppedSlice = removeElement(originalSlice, 100)
	fmt.Printf("original slice: %v, cropped slice: %v \n", originalSlice, croppedSlice)
	croppedSlice = removeElement(originalSlice, -20)
	fmt.Printf("original slice: %v, cropped slice: %v \n", originalSlice, croppedSlice)

}

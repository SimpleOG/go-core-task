package main

import "fmt"

func FoundEntries(s1, s2 []string) []string {
	entries := make([]string, 0, len(s1)/2)
	s2Map := make(map[string]struct{}, len(s2))
	//получим все значения которые есть в s2
	for _, v := range s2 {
		s2Map[v] = struct{}{}
	}
	//пробежимся по первому слайсу
	for _, v := range s1 {
		//если текущее значение из s1 отсутствует в s2, то добавляем в ответ
		if _, ok := s2Map[v]; !ok {
			entries = append(entries, v)
		}
	}
	return entries
}
func main() {
	//Напишите функцию, которая возвращает слайс строк,
	//содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := FoundEntries(slice1, slice2)
	fmt.Println(result)
}

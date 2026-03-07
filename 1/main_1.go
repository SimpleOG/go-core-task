package main

import (
	"fmt"
	"strconv"
)

// Сделал определение 1 конкретной переменной
func getType(v any) string {
	return fmt.Sprintf("%T", v)
}

// Вариативное кол-во переменных на вход
func typeToString(args ...any) string {
	s := ""
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			s += strconv.Itoa(v) + " "
		case float64:
			s += strconv.FormatFloat(v, 'f', -1, 64) + " "
		case string:
			s += v + " "
		case bool:
			s += strconv.FormatBool(v) + " "
		case complex64:
			s += fmt.Sprintf("%v", v) + " "
		}
	}
	return s
}
func main() {
	// 1. Создает несколько переменных различных типов данных:
	var (
		dec  int       = 42
		oct  int       = 052
		hex  int       = 0x2A
		flt  float64   = 2.71
		str  string    = "hello"
		flag bool      = true
		comp complex64 = 1 + 2i
	)

	// 2. Определяет тип каждой переменной и выводит его на экран.
	fmt.Println(getType(dec))
	fmt.Println(getType(oct))
	fmt.Println(getType(hex))
	fmt.Println(getType(flt))
	fmt.Println(getType(str))
	fmt.Println(getType(flag))
	fmt.Println(getType(comp))

	// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
	s := typeToString(dec, oct, hex, flt, str, flag, comp)
	fmt.Println(s)

	// 4. Преобразовать эту строку в срез рун.
	runeS := []rune(s)

	fmt.Println(runeS)

}

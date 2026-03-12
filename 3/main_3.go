package main

import "fmt"

type Pair struct {
	key   string
	value int
}
type StringIntMap struct {
	buckets [][]Pair
}

func NewCustomMap(mapSize int) CustomMap {
	return &StringIntMap{
		buckets: make([][]Pair, mapSize),
	}
}

type CustomMap interface {
	Add(key string, value int)
	Remove(key string)
	Copy() map[string]int
	Exists(key string) bool
	Get(key string) (int, bool)
}

// аналог хешфункции
func (m *StringIntMap) Hash(key string) int {
	hash := 0
	for _, v := range key {
		hash += int(v)
	}
	return hash % len(m.buckets)
}

// Добавление элемента: Метод Add(key string, value int),
// который добавляет новую пару "ключ-значение" в карту.
func (m *StringIntMap) Add(key string, value int) {
	//вычисляем хеш
	index := m.Hash(key)
	currentBucket := m.buckets[index]
	for i, pair := range currentBucket {
		//если ключ уже есть в мапе - перезаписываем
		if pair.key == key {
			m.buckets[index][i].value = value
			return
		}
	}
	pair := Pair{key: key, value: value}
	m.buckets[index] = append(m.buckets[index], pair)
}

// Удаление элемента: Метод Remove(key string),
// который удаляет элемент по ключу из карты.
func (m *StringIntMap) Remove(key string) {
	index := m.Hash(key)
	//Нам нужен только бакет в котором лежат хешы
	//совпадающие с хешем переданного ключа
	currentBucket := m.buckets[index]
	for i, pair := range currentBucket {
		if pair.key == key {
			m.buckets[index] = append(currentBucket[:i], currentBucket[i+1:]...)
			return
		}
	}
}

// Копирование карты: Метод Copy() map[string]int,
// который возвращает новую карту, содержащую все элементы текущей карты.
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(m.buckets))
	for _, bucket := range m.buckets {
		for _, pair := range bucket {
			newMap[pair.key] = pair.value
		}
	}
	return newMap
}

// Получение значения: Метод Get(key string) (int, bool),
// который возвращает значение по ключу и булевый флаг,
// указывающий на успешность операции.
func (m *StringIntMap) Get(key string) (int, bool) {
	index := m.Hash(key)
	bucket := m.buckets[index]
	for _, p := range bucket {
		if p.key == key {
			return p.value, true
		}
	}
	return 0, false
}

// Проверка наличия ключа: Метод Exists(key string) bool,
// который проверяет, существует ли ключ в карте.
func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.Get(key)
	return ok
}

func main() {
	m := NewCustomMap(8)
	m.Add("Джин", 33)
	m.Add("Шуга", 33)
	m.Add("Чонгук", 28)
	m.Add("Чимин", 30)
	m.Add("Ви", 30)
	m.Remove("Ви")
	age, ok := m.Get("Чимин")
	if ok {
		fmt.Printf("Возраст %s=%v \n", "Чимин", age)
	}
	exists := m.Exists("Чонгук")
	fmt.Printf("Чонгук exists?  %v\n", exists)
	newBTS := m.Copy()
	fmt.Println(newBTS)

}

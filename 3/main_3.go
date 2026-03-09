package main

type Pair struct {
	key   string
	value int
}
type StringIntMap struct {
	bucket  [][]Pair
	mapSize int
}
type CustomMap interface {
	Add(key string, value int)
	Remove(key string)
	Copy() map[string]int
	Exists(key string) bool
	Get(key string) (int, bool)
}

func (s *StringIntMap) Add(key string, value int) {
	pair := &Pair{key: key, value: value}

}
func (s *StringIntMap) Remove(key string) {

}
func (s *StringIntMap) Copy() map[string]int {

}
func (s *StringIntMap) Exists(key string) bool {

}
func (s *StringIntMap) Get(key string) (int, bool) {

}

func NewCustomMap(mapSize int) CustomMap {
	return &StringIntMap{
		bucket:  make([][]Pair, mapSize),
		mapSize: mapSize,
	}
}

//

func main() {

}

/*К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации. */

package main

// С этой задачкой без помощи чатика стажеров я не разобрался :(

// Тут проблема с память, слишком большая V создается в памяти и только часть копируется в justString
// Из за этого большая часть памяти выделенная под v остается неиспользуемой

func createHugeString(size uint64) string {
	return string(make([]rune, size))
}

func makeSlice(v string, from, to uint64) string {
	data := []rune(v)
	slice := make([]rune, 0)
	for i := from; i < to; i++ {
		slice = append(slice, data[i])
	}
	return string(slice)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = makeSlice(v, 0, 100)
}

func main() {
	someFunc()
}

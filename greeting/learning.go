package greeting

import (
	"strconv"
)

func Teacharr() []int {
	arr := []int{}       // это массив
	arr = append(arr, 1) // добавление в массив
	return arr
}

func Teachmath() int {
	a := 5
	a += 1
	a -= 1
	a /= 5
	a *= 5
	a %= 2
	return a
}

func Teachfor() []string {
	names := []string{"John", "Harold", "Vince"}
	// i — это индекс, name — это значение на текущем шаге цикла
	greet_arr := []string{}
	for i, name := range names {
		greet_arr = append(greet_arr, name, strconv.Itoa(i))
	}
	return greet_arr
}

func Teachmap() map[int]string {
	m := map[int]string{1: "hello", 2: "world"}
	m[3] = "!"
	return m
}

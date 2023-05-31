package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hexlet-go/greeting"
)

func main() {
	f := greeting.Hello()
	fmt.Println(color.GreenString(f))
	arr := greeting.Teacharr()
	fmt.Println(arr)
	math := greeting.Teachmath()
	fmt.Println(math)
	ff := greeting.Teachfor()
	fmt.Println("массив:", ff)
	fmt.Println("Длина массива", len(ff))
	mm := greeting.Teachmap()
	fmt.Println("Хеш таблица, словарь:", mm)
	fmt.Println("Длина cловаря", len(mm))
}

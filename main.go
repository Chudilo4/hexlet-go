package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hexlet-go/greeting"
)

func main() {
	f := greeting.Hello()
	fmt.Println(color.GreenString(f))
}

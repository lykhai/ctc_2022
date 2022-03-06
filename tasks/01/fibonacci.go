package main

import (
	"fmt"
)

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	temp := 0

	return func() int {
		temp = f0 + f1
		f0 = f1
		f1 = temp
		return f0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

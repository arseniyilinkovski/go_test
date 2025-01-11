package main

import (
	"fmt"
	"os"
)

func main() {
	var a float32
	var b float32
	var choice int16
	fmt.Print("Введите 1 число: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Print("Введите 2 число: ")
	fmt.Fscan(os.Stdin, &b)
	fmt.Print("Выберите действие: \n1) +\n2) -\n3) *\n4)/\n")
	fmt.Fscan(os.Stdin, &choice)
	switch choice {
	case 1:
		fmt.Print("Результат: ", get_sum(a, b))
	case 2:
		fmt.Print("Результат: ", get_minus(a, b))
	case 3:
		fmt.Print("Результат: ", get_mul(a, b))
	case 4:
		fmt.Print("Результат: ", get_div(a, b))
	}
}

func get_sum(a float32, b float32) float32 {
	return a + b
}

func get_minus(a float32, b float32) float32 {
	return a - b
}

func get_mul(a float32, b float32) float32 {
	return a * b
}

func get_div(a float32, b float32) float32 {
	return a / b
}

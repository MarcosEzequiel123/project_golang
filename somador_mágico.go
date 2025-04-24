package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var op string

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&a)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&b)

	fmt.Println("Digite a operação (+, -, *, /):")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Println("Resultado:", a+b)
	case "-":
		fmt.Println("Resultado:", a-b)
	case "*":
		fmt.Println("Resultado:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Resultado:", a/b)
		} else {
			fmt.Println("Erro: divisão por zero!")
		}
	default:
		fmt.Println("Operação inválida.")
	}
}

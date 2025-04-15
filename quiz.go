package main

import (
	"fmt"
	"strings"
)

func main() {
	// Pergunta 1
	fmt.Println("Quem é Saitama?")
	fmt.Println("a) Um herói de Dragon Ball")
	fmt.Println("b) Um herói de One Punch Man")
	fmt.Println("c) Um vilão de Naruto")
	fmt.Println("d) Um super-herói da Marvel")

	var resposta1 string
	fmt.Scanln(&resposta1)

	if strings.ToLower(resposta1) == "b" {
		fmt.Println("Resposta correta! Saitama é o herói de One Punch Man.")
	} else {
		fmt.Println("Resposta incorreta. A resposta correta é 'b'.")
	}

	// Pergunta 2
	fmt.Println("\nQual é o poder principal de Saitama?")
	fmt.Println("a) Super velocidade")
	fmt.Println("b) Força extrema com um soco")
	fmt.Println("c) Habilidade de voar")
	fmt.Println("d) Telepatia")

	var resposta2 string
	fmt.Scanln(&resposta2)

	if strings.ToLower(resposta2) == "b" {
		fmt.Println("Resposta correta! O poder principal de Saitama é a força extrema com um soco.")
	} else {
		fmt.Println("Resposta incorreta. A resposta correta é 'b'.")
	}

	// Pergunta 3
	fmt.Println("\nQual é o nome do vilão principal de One Punch Man no início da série?")
	fmt.Println("a) Boros")
	fmt.Println("b) Garou")
	fmt.Println("c) Dr. Kuseno")
	fmt.Println("d) Crablante")

	var resposta3 string
	fmt.Scanln(&resposta3)

	if strings.ToLower(resposta3) == "d" {
		fmt.Println("Resposta correta! O vilão principal no início é Crablante.")
	} else {
		fmt.Println("Resposta incorreta. A resposta correta é 'd'.")
	}
}

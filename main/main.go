package main

import (
	"fmt"
	"os"
	"strconv"

	calculo "github.com/timoteoBone/goGithubTry/calulator"
)

func main() {
	nums := []int{}
	for _, num := range os.Args[2:] {
		aux, err := strconv.Atoi(num)
		if err == nil {
			nums = append(nums, aux)
		}
	}

	result, err := operate(os.Args[1], nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func operate(operador string, numeros []int) (int, error) {
	var error error
	var result int
	switch operador {
	case "+":
		result = calculo.Sumar(numeros)
	case "-":
		result = calculo.Restar(numeros)
	case "/":
		result, error = calculo.Dividir(numeros[0], numeros[1])
	case "++":
		result = calculo.Multiplicar(numeros)
	}

	return result, error
}

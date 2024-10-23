package main

import (
	"fmt"
	"proj1/structs"
)

const eboulicaoK = 373.15

func main() {

	temp := structs.Temperatura{
		Kelvin: eboulicaoK,
	}

	var err error
	temp.Celsius, err = structs.ConverterKelvinParaCelsius(temp.Kelvin)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("A temperatura de %v°K é equivalente a %v°C\n", temp.Kelvin, temp.Celsius)
}

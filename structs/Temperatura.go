package structs

import "errors"

type Temperatura struct {
	Kelvin  float64
	Celsius float64
}

func ConverterKelvinParaCelsius(kelvin float64) (float64, error) {
	if kelvin < 0 {
		return 0, errors.New("a temperatura em Kelvin não pode ser negativa")
	}
	return kelvin - 273.15, nil
}

package entity

type InputCep struct {
	Cep string `json:"cep"`
}

type ClimaOutput struct {
	City       string  `json:"city"`
	Celsius    float32 `json:"temp_C"`
	Fahrenheit float32 `json:"temp_F"`
	Kelvin     float32 `json:"temp_K"`
}

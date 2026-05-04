package entity

type Clima struct {
	Location Location
	Current  Current
}

type Location struct {
	Name            string  `json:"name"`
	Region          string  `json:"region"`
	Country         string  `json:"country"`
	Lat             float64 `json:"lat"`
	Lon             float64 `json:"lon"`
	Tz_id           string  `json:"tz_id"`
	Localtime_epoch int32   `json:"localtime_epoch"`
	Localtime       string  `json:"localtime"`
}

type Current struct {
	Last_updated_epoch int32   `json:"last_updated_epoch"`
	Last_updated       string  `json:"last_updated"`
	Temp_c             float32 `json:"temp_c"`
	Temp_f             float32 `json:"temp_f"`
	Is_day             int32   `json:"is_day"`
}

type ClimaOutput struct {
	City       string  `json:"city"`
	Celsius    float32 `json:"temp_C"`
	Fahrenheit float32 `json:"temp_F"`
	Kelvin     float32 `json:"temp_K"`
}

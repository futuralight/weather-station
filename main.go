package main

import (
	"fmt"

	"github.com/futuralight/weather-station/sensor"
)

func main() {
	sensor, err := sensor.CreateBme280Sensor(1)
	if err != nil {
		panic(err)
	}
	t, err := sensor.SenseTemperature()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Температура %f", t))
	p, err := sensor.SensePressure()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Давление %f", p))
	h, err := sensor.SenseHumidity()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Влажность %f", h))
}

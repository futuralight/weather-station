package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/futuralight/weather-station/sensor"
	"github.com/joho/godotenv"
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

func init() {
	err := loadEnv()
	if err != nil {
		panic(err)
	}
}

func loadEnv() error {
	godotenv.Load()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	err = godotenv.Load(dir + "/.env") //Загрузка .env файла
	return nil
}

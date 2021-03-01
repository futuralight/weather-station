package main

import (
	"fmt"

	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
)

func main() {
	defer logger.FinalizeLogger()
	// Create new connection to i2c-bus on 1 line with address 0x76.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2c, err := i2c.NewI2C(0x76, 1)
	if err != nil {
		panic(err)
	}
	defer i2c.Close()
	// Uncomment/comment next lines to suppress/increase verbosity of output
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c) // signature=0x60
	if err != nil {
		panic(err)

	}
	id, err := sensor.ReadSensorID()
	if err != nil {
		panic(err)

	}
	fmt.Println(id)
	err = sensor.IsValidCoefficients()
	if err != nil {
		panic(err)

	}
	// Read temperature in celsius degree
	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	fmt.Println(fmt.Sprintf("Температура %fС", t))

	supported, h1, err := sensor.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
	if supported {
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Влажность = %v %%", h1))
	}
	p, err := sensor.ReadPressureMmHg(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Давление = %f %%", p))
}

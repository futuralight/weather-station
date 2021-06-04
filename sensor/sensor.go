package sensor

import (
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
)

//Sensitive - sensor interface
type Sensitive interface {
	SenseTemperature() (float32, error)
	SensePressure() (float32, error)
	SenseHumidity() (float32, error)
}

//BME280Sensor - 280 wrapper struct
type BME280Sensor struct {
	sensor *bsbmp.BMP
}

//SenseTemperature - get bme280 temperature
func (bme280 *BME280Sensor) SenseTemperature() (float32, error) {
	return bme280.sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
}

//SensePressure get 280 pressure
func (bme280 *BME280Sensor) SensePressure() (float32, error) {
	return bme280.sensor.ReadPressureMmHg(bsbmp.ACCURACY_STANDARD)
}

//SenseHumidity get 280 humidity
func (bme280 *BME280Sensor) SenseHumidity() (float32, error) {
	_, hmd, err := bme280.sensor.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
	return hmd, err
}

//CreateBme280Sensor - return brand new BME280Sensor
func CreateBme280Sensor(bus int) (*BME280Sensor, error) {
	defer logger.FinalizeLogger()
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel) //TODO убрать
	logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)
	i2c, err := i2c.NewI2C(0x76, bus)
	if err != nil {
		return nil, err
	}
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c)
	if err != nil {
		return nil, err
	}
	return &BME280Sensor{
		sensor,
	}, nil
}

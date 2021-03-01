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

//BME280Sensor - 280 wrapper
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

//SenseHumidity get 280 Humidity
func (bme280 *BME280Sensor) SenseHumidity() (float32, error) {
	_, hmd, err := bme280.sensor.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
	return hmd, err
}

//CreateBme280Sensor - return brand new BME280Sensor
func CreateBme280Sensor(bus int) (BME280Sensor, error) {

	defer logger.FinalizeLogger()
	// Uncomment/comment next lines to suppress/increase verbosity of output
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)
	var bme280Sensor BME280Sensor
	i2c, err := i2c.NewI2C(0x76, bus)
	if err != nil {
		return bme280Sensor, err
	}
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c) // signature=0x60
	if err != nil {
		return bme280Sensor, err
	}
	bme280Sensor.sensor = sensor
	return bme280Sensor, nil
}

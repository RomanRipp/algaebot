package sensors

import (
	"github.com/stianeikeland/go-rpio"
)

type Light struct {
	pinNumber int
}

func (l Light) Read() (float64, error) {
	err := rpio.Open()
	reading := 0.0
	if  err == nil {
		pin := rpio.Pin(l.pinNumber)
		pin.Input()
		switch pin.Read() {
		case 0:
			reading = 0.0
		case 1:
			reading = 1.0
		}
	}
	defer rpio.Close()
	return reading, err
}

func (l Light) IsUnderLight() (bool, error) {
	value, error := l.Read()
	return value == 0, error
}

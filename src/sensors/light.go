package sensors

import (
	"github.com/stianeikeland/go-rpio"
)

type Light struct {
	pinNumber int
}

func (s Light) Read() (float64, error) {
	err := rpio.Open()
	reading := 0.0
	if  err == nil {
		pin := rpio.Pin(s.pinNumber)
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

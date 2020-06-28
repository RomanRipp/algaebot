package sensors

import (
	"github.com/stianeikeland/go-rpio"
)

type Light struct {
	pinNumber int
	lastReading float64
}

func (s Light) Read() (float64, error) {
	err := rpio.Open()
	if  err == nil {
		pin := rpio.Pin(s.pinNumber)
		pin.Input()
		switch pin.Read() {
		case 0:
			s.lastReading = 0.0
		case 1:
			s.lastReading = 1.0
		}
	}
	defer rpio.Close()
	return s.lastReading, err
}

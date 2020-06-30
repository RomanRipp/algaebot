package sensors

import "github.com/stianeikeland/go-rpio"

type Battery struct {
	batteryPinNum int
}

func (b *Battery) Info() {
	if b.batteryPinNum == 0 {
		b.batteryPinNum = 35
	}
}

func (b Battery) IsCharged() (bool, error) {
	err := rpio.Open()
	if err == nil {
		pin := rpio.Pin(b.batteryPinNum)
		pin.Input()
		return pin.Read() == rpio.High, err
	}
	defer rpio.Close()
	return false, err
}




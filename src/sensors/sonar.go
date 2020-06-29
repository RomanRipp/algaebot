package sensors

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type Sonar struct {
	triggerPinNum int
	echoPinNum    int
}

func (s Sonar) ReadDistance() (float64, error) {
	err := rpio.Open()
	distance := 0.0
	if err == nil {
		echoPin := rpio.Pin(s.echoPinNum)
		echoPin.Mode(rpio.Input)
		echoPin.PullDown()

		triggerPin := rpio.Pin(s.triggerPinNum)
		triggerPin.Mode(rpio.Output)
		triggerPin.Low()

		// Wait for sensor to settle.
		time.Sleep(time.Second)
		triggerPin.High()
		time.Sleep(time.Microsecond)
		triggerPin.Low()

		pulseStart := time.Now()
		for echoPin.Read() == 0 {
			pulseStart = time.Now()
		}

		pulseEnd := time.Now()
		for echoPin.Read() == 0 {
			pulseEnd = time.Now()
		}

		pulseDuration := pulseEnd.Nanosecond() - pulseStart.Nanosecond()
		distance = float64(pulseDuration * 17150.0)
	}
	defer rpio.Close()
	return distance, err
}


package sensors

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type Sonar struct {
	triggerPinNum int
	echoPinNum    int
}

const speedOfSoundInMmNs float64 = 343 * 1e-6

func (s Sonar) ReadDistance() (float64, error) {
	err := rpio.Open()
	distance := 0.0
	if err == nil {
		echoPin := rpio.Pin(s.echoPinNum)
		echoPin.Input()
		echoPin.PullDown()

		triggerPin := rpio.Pin(s.triggerPinNum)
		triggerPin.Output()
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
		for echoPin.Read() == 1 {
			pulseEnd = time.Now()
		}

		pulseDuration := pulseEnd.Sub(pulseStart)
		distance = (speedOfSoundInMmNs * float64(pulseDuration.Nanoseconds())) / 2
	}
	defer rpio.Close()
	return distance, err
}


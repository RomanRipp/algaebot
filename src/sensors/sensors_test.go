package sensors

import (
	"fmt"
	"testing"
	"time"
)

func TestLightSensor(t *testing.T) {
	light := Light{}
	light.pinNumber = 16
	value, error := light.Read()
	if value !=  0.0 || error != nil {
		t.Error("Invalid sensor reading")
	}
}

func TestSonarSensor(t *testing.T) {
	sonar := Sonar{triggerPinNum: 36, echoPinNum: 31}
	for {
		distance, error := sonar.ReadDistance()
		if error == nil {
			fmt.Print(distance)
			time.Sleep(time.Second)
		} else {
			t.Error("Failed to read sensor.")
		}
	}
}
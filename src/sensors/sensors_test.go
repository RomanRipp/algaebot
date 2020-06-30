package sensors

import (
	"testing"
)

func TestLightSensor(t *testing.T) {
	light := Light{}
	light.pinNumber = 23
	value, error := light.Read()
	if value !=  0.0 || error != nil {
		t.Error("Invalid sensor reading")
	}
}

func TestSonarSensor(t *testing.T) {
	sonar := Sonar{triggerPinNum: 16, echoPinNum: 6}
	distance, error := sonar.ReadDistance()
	if error == nil {
		t.Log(distance)
	} else {
		t.Error("Failed to read sensor.")
	}
}
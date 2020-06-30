package sensors

import (
	"testing"
)

func TestLightSensor(t *testing.T) {
	light := Light{}
	light.pinNumber = 23
	value, error := light.Read()
	if error != nil {
		t.Error("Invalid sensor reading")
	} else {
		t.Log("Photo sensor reading: ", value)
	}

}

func TestSonarSensor(t *testing.T) {
	sonar := Sonar{triggerPinNum: 16, echoPinNum: 6}
	distance, error := sonar.ReadDistance()
	if error == nil {
		t.Log("Distance sensor reading: ", distance)
	} else {
		t.Error("Failed to read sensor.")
	}
}
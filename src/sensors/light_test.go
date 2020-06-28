package sensors

import "testing"

func TestLightSensor(t *testing.T) {
	light := Light{}
	light.pinNumber = 16
	value, error := light.Read()
	if value !=  0.0 || error != nil {
		t.Error("Invalid sensor reading")
	}
}

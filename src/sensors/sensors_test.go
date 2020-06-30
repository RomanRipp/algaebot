package sensors

import (
	"testing"
)

func TestLight_IsUnderLight(t *testing.T) {
	light := Light{}
	light.pinNumber = 23
	value, error := light.IsUnderLight()
	if error != nil {
		t.Error("Invalid sensor reading")
	} else {
		t.Log("Photo sensor reading: ", value)
	}

}

func TestSonar_ReadDistance(t *testing.T) {
	sonar := Sonar{triggerPinNum: 16, echoPinNum: 6}
	distance, error := sonar.ReadDistance()
	if error == nil {
		t.Log("Distance sensor reading: ", distance)
	} else {
		t.Error("Failed to read sensor.")
	}
}

func TestBattery_IsCharged(t *testing.T) {
	battery := Battery{}
	charged, err := battery.IsCharged()
	if err == nil && charged {
		t.Log("Battery is charged: ", charged)
	} else {
		t.Error("Failed to check battery charge.")
	}
}
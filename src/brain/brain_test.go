package brain

import "testing"

func TestFullBattery(t *testing.T) {
	_, ok := solve(robotState{charge: true}).(stay)
	if !ok {
		t.Error("Should stay if charge charged")
	}
}

func TestTooDarkAndTooBright(t *testing.T) {
	state := robotState{charge: false, light: false}
	_, ok := solve(state).(stay)
	if !ok {
		t.Error("Should stay if it is too dark.")
	}

	state = robotState{charge: false, light: false}
	_, ok = solve(state).(stay)
	if !ok {
		t.Error("Should stay if it is too dark.")
	}
}

func TestEdge(t *testing.T) {
	state := robotState{charge: false, light: true, edge: true}
	action, ok := solve(state).(drive)
	if !ok {
		t.Error("Should drive back if on the edge.")
	}
	if action.distance != step {
		t.Error("Incorrect distance chosen.")
	}
	if action.forward {
		t.Error("Incorrect drive direction.")
	}
}

func TestObstacleAhead(t *testing.T) {
	state := robotState{charge: false, light: true, edge: false, sonar: step - 1}
	action, ok := solve(state).(rotate)
	if !ok {
		t.Error("Incorrect action.")
	}
	if action.degree != rotation {
		t.Error("Incorrect rotation.")
	}
}

func TestMoveForward(t *testing.T) {
	state := robotState{charge: false, light: true, edge: false, sonar: step + 1}
	action, ok := solve(state).(drive)
	if !ok {
		t.Error("Incorrect action.")
	}
	if action.distance != step || !action.forward {
		t.Error("Incorrect motion.")
	}
}

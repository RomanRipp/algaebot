package brain

import "testing"

func TestFullBattery(t *testing.T) {
	_, ok := solve(robotState{battery: charged + 1}).(stay)
	if !ok {
		t.Error("Should stay if battery charged")
	}
}

func TestTooDarkAndTooBright(t *testing.T) {
	state := robotState{battery: charged - 1, light: tooDark - 1}
	_, ok := solve(state).(stay)
	if !ok {
		t.Error("Should stay if it is too dark.")
	}

	state = robotState{battery: charged - 1, light: tooBright + 1}
	_, ok = solve(state).(stay)
	if !ok {
		t.Error("Should stay if it is too dark.")
	}
}

func TestEdge(t *testing.T) {
	state := robotState{battery: charged - 1, light: tooBright - 1, edge: true}
	action, ok := solve(state).(drive)
	if !ok {
		t.Error("Should drive back if on the edge.")
	}
	if action.distance != step {
		t.Error("Incorrect distance shosen.")
	}
	if action.forward {
		t.Error("Incorect drive direction.")
	}
}

func TestObstacleAhead(t *testing.T) {
	state := robotState{battery: charged - 1, light: tooBright - 1, edge: false, sonar: step - 1}
	action, ok := solve(state).(rotate)
	if !ok {
		t.Error("Incorrect action.")
	}
	if action.degree != rotation || !action.right {
		t.Error("Incorrect rotation.")
	}
}

func TestMoveForward(t *testing.T) {
	state := robotState{battery: charged - 1, light: tooBright - 1, edge: false, sonar: step + 1}
	action, ok := solve(state).(drive)
	if !ok {
		t.Error("Incorrect action.")
	}
	if action.distance != step || !action.forward {
		t.Error("Incorrect motion.")
	}
}

package brain

var isRunning = true

type robotState struct {
	battery    float32
	light      float32
	sonar      float32
	edge       bool
	discovered [][]int
}

// Run starts the robot
func Run() {
	for isRunning {
		// light := light.RetrieveReading()
		// battery := float32(0.0)
		// state := robotState{light, battery, nil}
		// think(state)
	}
}

type action interface {
	execute() error
}

type drive struct {
	distance int
	forward  bool
}

func (d drive) execute() error {
	return nil
}

type rotate struct {
	degree int
	right  bool
}

func (d rotate) execute() error {
	return nil
}

type stay struct {
}

func (d stay) execute() error {
	return nil
}

const charged = 50
const tooBright = 90
const tooDark = 10
const step = 100
const rotation = 90

func solve(state robotState) action {
	if state.battery < charged {
		if tooDark < state.light && state.light < tooBright {
			if state.edge {
				return drive{step, false}
			}
			if state.sonar > step {
				return drive{step, true}
			}
			return rotate{rotation, true}
		}
	}
	return stay{}
}

// Stop closes the main loop.
func Stop() {
	isRunning = false
}

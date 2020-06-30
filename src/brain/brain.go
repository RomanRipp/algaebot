package brain

import (
	"log"
	"math/rand"
	"sensors"
)

type Brain struct {
	isRunning bool
	slam [][]bool
	trail []action
}

func NewBrain() *Brain {
	brain := new(Brain)
	brain.isRunning = false
	return brain
}

// Run starts the robot
func (b Brain) Run() {
	for b.isRunning {
		state := retrieveState()
		b.makeStep(state)
	}
}

// Stop closes the main loop.
func (b Brain) Stop() {
	b.isRunning = false
}

func (b Brain) makeStep(state robotState) {
	if state.error == nil {
		action := solve(state)
		action.execute()
		b.trail = append(b.trail, action)
		b.update(state, action)
	} else {
		log.Fatal("Failed to read sensors.")
	}
}

func (b Brain) update(state robotState, action action) {
	if len(b.slam) == 0 {
		row := make([]bool, 1)
		row[0] = state.light
		b.slam = append(b.slam, row)
	}
}

type robotState struct {
	charge bool // true if battery charged
	light  bool // true if under light
	sonar  float64 // mm of space before robot
	edge   bool // true if robot on the edge
	error error
}

func retrieveState() robotState {
	state := robotState{}
	lightSensor := sensors.Light{}
	if state.light, state.error = lightSensor.IsUnderLight(); state.error == nil {
		battery := sensors.Battery{}
		if state.charge, state.error = battery.IsCharged(); state.error == nil {
			sonar := sensors.Sonar{}
			if state.sonar, state.error = sonar.ReadDistance(); state.error == nil {
				state.edge = false
			}
		}
	}
	return state
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

const step = 100
const rotation = 90

func random() bool {
	return rand.Intn(2) == 0
}

func solve(state robotState) action {
	if !state.charge {
		if state.light {
			if state.edge {
				return drive{step, false}
			}
			if state.sonar > step {
				return drive{step, true}
			}
			return rotate{rotation, random()}
		}
	}
	return stay{}
}


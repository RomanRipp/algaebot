package sensors

type Battery struct {

}

func (s Battery) ReadCharge() (float64, error) {
	return 0.0, nil
}




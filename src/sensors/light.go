package sensors

type Light struct {

}

func (s Light) Read() (float64, error) {
	return 0.0, nil
}

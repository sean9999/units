package units

// a Magnitude unambiguously indicates a magnitude of something by combining a scalar and a unit
// example: 3.5 teaspons
type Magnitude struct {
	Unit     Unit
	Quantity float64
}

func (from Magnitude) Convert(to Unit) (float64, error) {
	ratio, err := from.Unit.Convert(to)
	if err != nil {
		return 0, err
	}
	return (ratio * from.Quantity), nil
}

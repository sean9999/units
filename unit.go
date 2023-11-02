package units

// a UnitKind indicates what kind of unit we're dealing with
// and can be used to enforce legal conversions.
// for example, it's not legal to convert between mile and teaspoon
type UnitKind uint8

const (
	Volume UnitKind = iota
	Weight
	Length
)

// a Unit is a unit of measure
type Unit struct {
	Symbol string
	Name   string
	Kind   UnitKind
	Ratio  float64
}

// UnitMismatchError is an [Error] that contains information about the mismatched units
type UnitMismatchError struct {
	Unit1 Unit
	Unit2 Unit
	Msg   string
}

func (e UnitMismatchError) Error() string {
	return e.Msg
}

// to Convert between two units is to return their Ratio
func (from Unit) Convert(to Unit) (float64, error) {
	if from.Kind != to.Kind {
		return 0, UnitMismatchError{from, to, "Unit Mismatch"}
	}
	return (from.Ratio / to.Ratio), nil
}

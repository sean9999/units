package volume

import "github.com/sean9999/units"

var Liter units.Unit = units.Unit{
	Kind:   units.Volume,
	Name:   "Liter",
	Symbol: "L",
	Ratio:  1,
}

var Milliliter units.Unit = units.Unit{
	Kind:   units.Volume,
	Name:   "Millileter",
	Symbol: "ml",
	Ratio:  1000,
}

var Teaspoon units.Unit = units.Unit{
	Kind:   units.Volume,
	Name:   "Teaspoon",
	Symbol: "tsp",
	Ratio:  4.93 / 1000,
}

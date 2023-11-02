package main

import (
	"fmt"

	"github.com/sean9999/units"
	"github.com/sean9999/units/volume"
)

func main() {

	fiveTeaspoons := units.Magnitude{
		Unit:     volume.Teaspoon,
		Quantity: 5,
	}

	fiveTeaspoonsInLiters, err := fiveTeaspoons.Convert(volume.Liter)
	if err != nil {
		panic(err)
	}

	fmt.Printf("five teaspoons is %f liters", fiveTeaspoonsInLiters)

}

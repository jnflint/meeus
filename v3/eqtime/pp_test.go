// Copyright 2013 Sonia Keys
// License: MIT

//go:build !nopp
// +build !nopp

package eqtime_test

import (
	"fmt"

	"github.com/jnflint/meeus/v3/eqtime"
	"github.com/jnflint/meeus/v3/julian"
	pp "github.com/jnflint/meeus/v3/planetposition"
	sexa "github.com/soniakeys/sexagesimal"
)

func ExampleE() {
	// Example 28.a, p. 184
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	j := julian.CalendarGregorianToJD(1992, 10, 13)
	eq := eqtime.E(j, earth)
	fmt.Printf("%+.1d", sexa.FmtHourAngle(eq))
	// Output:
	// +13ᵐ42ˢ.6
}

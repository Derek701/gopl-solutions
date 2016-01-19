// Package lenconv performs length conversions for feet and meters.
package lenconv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }

// FToM converts length in feet to meters.
func FToM(f Feet) Meters { return Meters(f * 0.3048) }

// MToF converts length in meters to feet.
func MToF(m Meters) Feet { return Feet(m / 0.3048) }

package clubs

import (
	"fmt"

	"github.com/rumere/clashcalc/pkg/clubdb"
)

type Category float64

const (
	Driver    Category = 1.0
	Wood               = 1.0
	LongIron           = 1.0
	ShortIron          = 1.0
	Wedge              = 1.0
	RoughIron          = 1.45
	SandWedge          = 1.15
)

type Correction float64

const (
	Grizzly Correction = 0.9
	B52                = 0.9
	Other              = 1.0
)

type Club struct {
	Category
	Correction
	Name      string
	Level     int64
	Power     int64
	Accuracy  int64
	TopSpin   int64
	BackSpin  int64
	Curl      int64
	BallGuide float64
}

func (c Club) SetCategory(s string) error {
	switch s {
	case "Driver":
		c.Category = Driver
	case "Wood":
		c.Category = Wood
	case "LongIron":
		c.Category = LongIron
	case "ShortIron":
		c.Category = ShortIron
	case "Wedge":
		c.Category = Wedge
	case "RoughIron":
		c.Category = RoughIron
	case "SandWedge":
		c.Category = SandWedge
	default:
		return fmt.Errorf("ERROR: Could not determine club category for %s.", s)
	}

	return nil
}

func (c Club) SetCorrection() {
	if c.Level >= 5 {
		if c.Name == "B52" {
			c.Correction = B52
		} else if c.Name == "Grizzly" {
			c.Correction = Grizzly
		}
	} else {
		c.Correction = Other
	}
}

func (c Club) GetClub(cdb clubdb.ClubData) error {
	idx := slices.IndexFunc(cdb, func(x clubdb.C) bool { return c.Key == "key1" })
}

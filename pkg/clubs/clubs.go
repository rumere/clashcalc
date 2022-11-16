/*
 * MIT License
 *
 * Copyright (c) 2018 golf-clash-notebook
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// clubs is a package for selecting and building the attributes for clubs
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

func (c Club) GetClub(cdb clubdb.ClubData, n string, l int64) error {
	found := false
	for i := range cdb {
		if cdb[i].Name == n && cdb[i].Level == l {
			found = true
			c.Accuracy = cdb[i].Accuracy
			c.BackSpin = cdb[i].BackSpin
			c.BallGuide = cdb[i].BallGuide
			c.Curl = cdb[i].Curl
			c.Level = cdb[i].Level
			c.Name = cdb[i].Name
			c.Power = cdb[i].Power
			c.TopSpin = cdb[i].TopSpin
			c.SetCategory(cdb[i].Category)
			c.SetCorrection()
		}
	}

	if found {
		return nil
	} else {
		return fmt.Errorf("ERROR: Unable to load club %s:%d", n, l)
	}
}

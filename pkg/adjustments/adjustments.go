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

// adjustments is a package to calculate the number of rings to move the target reletive to the wind given all inputs
//  (Club, Ball, Wind, Elevation, Slider)
package adjustments

import (
	"github.com/rumere/clashcalc/pkg/balls"
	"github.com/rumere/clashcalc/pkg/clubs"
	"github.com/rumere/clashcalc/pkg/wind"
)

// Adjustment is a struct to hold the all of the inputs needed to calculate the rings for an adjustment
type Adjustment struct {
	Club      clubs.Club
	Ball      balls.Ball
	Wind      wind.Wind
	Elevation int64
	Slider    int64
}

// Setter skeletons in case some logic is needed
func (a *Adjustment) SetClub(c clubs.Club) {
	a.Club = c
}
func (a *Adjustment) SetBall(b balls.Ball) {
	a.Ball = b
}
func (a *Adjustment) SetWind(w wind.Wind) {
	a.Wind = w
}
func (a *Adjustment) SetElevation(e int64) {
	a.Elevation = e
}
func (a *Adjustment) SetSlider(s int64) {
	a.Slider = s
}

// WindPerRing calculates the amount of wind for each ring given the inputs in the Adjustment struct
func (a *Adjustment) WindPerRing() error {
	return nil
}

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

// wind is a package that deals with wind inputs to be used in the adjustments package
package wind

// WindDirection is a type to hold constants for adjustments based on head/tail wind.
type WindDirection float64

// These values are swags.  They should work OK.  I imagine there is a fairly comp;lpicated formula to get more exact
//  numbers.
const (
	None WindDirection = 1.0
	N                  = 1.1
	NNE                = 1.09
	NE                 = 1.07
	ENE                = 1.04
	E                  = 1.0
	ESE                = 1.08
	SE                 = 1.14
	SSE                = 1.19
	S                  = 1.2
	SSW                = 1.19
	SW                 = 1.14
	WSW                = 1.08
	W                  = 1.0
	WNW                = 1.04
	NW                 = 1.07
	NNW                = 1.09
)

// WindDirection.String() converts the wind direction iota into printable/comparable short strings.  Because the
//  adjustments for mirrored wind directions are the same, it cannot be determined which exact wind direction to
//  return.  A string with both wind directions is returned.
func (wd WindDirection) String() string {
	switch wd {
	case N:
		return "N"
	case NNE:
		return "NNE or NNW"
	case NE:
		return "NE or NW"
	case ENE:
		return "ENE or WNW"
	case E:
		return "E or W"
	case ESE:
		return "ESE or WSW"
	case SE:
		return "SE or SW"
	case SSE:
		return "SSE or SSW"
	case S:
		return "S"
	}

	return "None"
}

// Wind is a struct to hold the wind inputs
type Wind struct {
	Speed float64
	WindDirection
}

// SetSpeed is a helper to set the wind speed
func (w *Wind) SetSpeed(s float64) {
	w.Speed = s
}

// SetWindDirection is a helper to set the wind direction value from a string
func (w *Wind) SetWindDirection(s string) {
	switch s {
	case "N":
		w.WindDirection = N
	case "NNE":
		w.WindDirection = NNE
	case "NE":
		w.WindDirection = NE
	case "ENE":
		w.WindDirection = ENE
	case "E":
		w.WindDirection = E
	case "ESE":
		w.WindDirection = ESE
	case "SE":
		w.WindDirection = SE
	case "SSE":
		w.WindDirection = SSE
	case "S":
		w.WindDirection = S
	case "SSW":
		w.WindDirection = SSW
	case "SW":
		w.WindDirection = SW
	case "WSW":
		w.WindDirection = WSW
	case "W":
		w.WindDirection = W
	case "WNW":
		w.WindDirection = WNW
	case "NW":
		w.WindDirection = NW
	case "NNW":
		w.WindDirection = NNW
	default:
		w.WindDirection = None
	}
}

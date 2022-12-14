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

// clubdb is a package for reading club data from persistence (currently a file)
package clubdb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var clubFilename = "clubdb.json"

type ClubData []struct {
	Name      string  `json:"name"`
	Level     int64   `json:"level"`
	Category  string  `json:"category"`
	Power     int64   `json:"power"`
	Accuracy  int64   `json:"accuracy"`
	TopSpin   int64   `json:"top-spin"`
	BackSpin  int64   `json:"back-spin"`
	Curl      int64   `json:"curl"`
	BallGuide float64 `json:"ball-guide"`
}

func (cd *ClubData) LoadData() error {
	jsonFile, err := os.Open(clubFilename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, cd)
	if err != nil {
		return err
	}

	return nil
}

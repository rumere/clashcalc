package clubdb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var clubFilename = "clubdb.json"

type ClubData []struct {
	Name      string  `json:"name"`
	Level     int     `json:"level"`
	Category  string  `json:"category"`
	Power     int     `json:"power"`
	Accuracy  int     `json:"accuracy"`
	TopSpin   int     `json:"top-spin"`
	BackSpin  int     `json:"back-spin"`
	Curl      int     `json:"curl"`
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

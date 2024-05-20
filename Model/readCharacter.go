package Model

import (
	"encoding/json"
	"io/ioutil"
)

func ReadCharacter(file string) (*PBCharacter, error) {
	var char PBCharacter
	charBytes, err := ioutil.ReadFile("./" + file)
	err = json.Unmarshal(charBytes, &char)
	if err != nil {
		return nil, err
	}
	return &char, nil
}

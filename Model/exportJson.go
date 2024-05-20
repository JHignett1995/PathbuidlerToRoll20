package Model

import (
	"encoding/json"
	"io/ioutil"
)

func ExportJson(generated *Roll20Character) error {
	jsonBytes, err := json.Marshal(generated)
	err = ioutil.WriteFile(generated.Character.Name+"_Converted.json", jsonBytes, 0777)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"PathbuidlerToRoll20/Model"
	"PathbuidlerToRoll20/constants"
	"io/ioutil"
	"strings"
)

func main() {

	dir, err := ioutil.ReadDir("./")
	constants.CheckErr(500, "Finding Characters", err)
	for _, char := range dir {
		if strings.Contains(char.Name(), ".json") {
			if !strings.Contains(char.Name(), "_Converted") {

				constants.CheckErr(200, "Starting Character: "+char.Name(), nil)
				PBcharacter, err := Model.ReadCharacter(char.Name())
				constants.CheckErr(500, "Result", err)

				constants.CheckErr(200, "Starting Conversion", nil)
				r20 := Model.Convert(PBcharacter)

				constants.CheckErr(200, "Writing File", nil)
				err = Model.ExportJson(r20)
				constants.CheckErr(500, "Output File", err)
				constants.CheckErr(200, "**************************Finished**************************", nil)
			}
		}
	}
}

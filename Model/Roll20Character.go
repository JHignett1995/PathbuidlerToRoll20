package Model

type Roll20Character struct {
	SchemaVersion int    `json:"schema_version"`
	Type          string `json:"type"`
	Character     struct {
		Name             string      `json:"name"`
		Bio              string      `json:"bio"`
		Avatar           string      `json:"avatar"`
		Gmnotes          string      `json:"gmnotes"`
		Defaulttoken     string      `json:"defaulttoken"`
		Tags             string      `json:"tags"`
		Controlledby     string      `json:"controlledby"`
		Inplayerjournals string      `json:"inplayerjournals"`
		Attribs          []Attribute `json:"attribs"`
		Abilities        []string    `json:"abilities"`
	} `json:"character"`
}

type Attribute struct {
	Name    string      `json:"name"`
	Current interface{} `json:"current"`
	Max     interface{} `json:"max"`
	ID      string      `json:"id"`
}

func NewAttribute(name string, current interface{}, ID string) *Attribute {
	return &Attribute{Name: name, Current: current, Max: "", ID: ID}
}
func NewAttribute1(name string, current interface{}, max interface{}, ID string) *Attribute {
	return &Attribute{Name: name, Current: current, Max: max, ID: ID}
}

func NewRoll20Character() *Roll20Character {
	return &Roll20Character{}
}

package Model

type SpellSearch struct {
	Count   int     `json:"count"`
	Results []Spell `json:"results"`
}

type Spell struct {
	ID   string `json:"_id"`
	Data struct {
		Ability struct {
			Value string `json:"value"`
		} `json:"ability"`
		Area struct {
			AreaType string      `json:"areaType"`
			Value    interface{} `json:"value"`
		} `json:"area"`
		Areasize struct {
			Value string `json:"value"`
		} `json:"areasize"`
		AutoHeightenLevel struct {
			Value interface{} `json:"value"`
		} `json:"autoHeightenLevel"`
		Category struct {
			Value string `json:"value"`
		} `json:"category"`
		Components struct {
			Focus    bool `json:"focus"`
			Material bool `json:"material"`
			Somatic  bool `json:"somatic"`
			Verbal   bool `json:"verbal"`
		} `json:"components"`
		Cost struct {
			Value string `json:"value"`
		} `json:"cost"`
		Damage struct {
			Value struct {
				Num0 struct {
					ApplyMod bool `json:"applyMod"`
					Type     struct {
						Categories []interface{} `json:"categories"`
						Subtype    interface{}   `json:"subtype"`
						Value      string        `json:"value"`
					} `json:"type,omitempty"`
					Value string `json:"value,omitempty"`
				} `json:"0,omitempty"`
			} `json:"value,omitempty"`
		} `json:"damage,omitempty"`
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
		Duration struct {
			Value string `json:"value"`
		} `json:"duration"`
		HasCounteractCheck struct {
			Value bool `json:"value"`
		} `json:"hasCounteractCheck"`
		Level struct {
			Value int `json:"value"`
		} `json:"level"`
		Location struct {
			Value interface{} `json:"value"`
		} `json:"location"`
		Materials struct {
			Value string `json:"value"`
		} `json:"materials"`
		Prepared struct {
			Value interface{} `json:"value"`
		} `json:"prepared"`
		Primarycheck struct {
			Value string `json:"value"`
		} `json:"primarycheck"`
		Range struct {
			Value string `json:"value"`
		} `json:"range"`
		Rules []interface{} `json:"rules"`
		Save  struct {
			Basic string `json:"basic"`
			Value string `json:"value"`
		} `json:"save"`
		Scaling struct {
			Damage struct {
				Num0 string `json:"0"`
			} `json:"damage"`
			Interval int `json:"interval"`
		} `json:"scaling"`
		School struct {
			Value string `json:"value"`
		} `json:"school"`
		Secondarycasters struct {
			Value string `json:"value"`
		} `json:"secondarycasters"`
		Secondarycheck struct {
			Value string `json:"value"`
		} `json:"secondarycheck"`
		Source struct {
			Value string `json:"value"`
		} `json:"source"`
		SpellType struct {
			Value string `json:"value"`
		} `json:"spellType"`
		Sustained struct {
			Value bool `json:"value"`
		} `json:"sustained"`
		Target struct {
			Value string `json:"value"`
		} `json:"target"`
		Time struct {
			Value string `json:"value"`
		} `json:"time"`
		Traditions struct {
			Value []string `json:"value"`
		} `json:"traditions"`
		Traits struct {
			Custom string `json:"custom"`
			Rarity struct {
				Value string `json:"value"`
			} `json:"rarity"`
			Value []string `json:"value"`
		} `json:"traits"`
	} `json:"data"`
	Effects []interface{} `json:"effects"`
	Name    string        `json:"name"`
	Type    string        `json:"type"`
}

func NewSpell(name string) *Spell {
	return &Spell{Name: name}
}

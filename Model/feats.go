package Model

type FeatSearch struct {
	Count   int       `json:"count"`
	Results []APIFeat `json:"results"`
}

type APIFeat struct {
	Id   string `json:"_id"`
	Data struct {
		ActionCategory struct {
			Value string `json:"value"`
		} `json:"actionCategory"`
		ActionType struct {
			Value string `json:"value"`
		} `json:"actionType"`
		Actions struct {
			Value *int `json:"value"`
		} `json:"actions"`
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
		FeatType struct {
			Value string `json:"value"`
		} `json:"featType"`
		Level struct {
			Value int `json:"value"`
		} `json:"level"`
		Prerequisites struct {
			Value []struct {
				Value string `json:"value"`
			} `json:"value"`
		} `json:"prerequisites"`
		Rules []struct {
			Key       string `json:"key"`
			Label     string `json:"label"`
			Predicate struct {
				All []string `json:"all"`
				Any []string `json:"any"`
			} `json:"predicate"`
			Selector string `json:"selector"`
			Type     string `json:"type"`
			Value    struct {
				Brackets []struct {
					End   int         `json:"end"`
					Start int         `json:"start"`
					Value interface{} `json:"value"`
				} `json:"brackets"`
				Field string `json:"field"`
			} `json:"value"`
		} `json:"rules"`
		Source struct {
			Value string `json:"value"`
		} `json:"source"`
		Traits struct {
			Custom string `json:"custom"`
			Rarity struct {
				Value string `json:"value"`
			} `json:"rarity"`
			Value []string `json:"value"`
		} `json:"traits"`
		Location string `json:"location,omitempty"`
	} `json:"data"`
	Effects []interface{} `json:"effects"`
	Name    string        `json:"name"`
	Type    string        `json:"type"`
}

func NewFeat(name string) *APIFeat {
	return &APIFeat{Name: name}
}

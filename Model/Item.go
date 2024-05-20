package Model

type ItemResp struct {
	Count   int    `json:"count"`
	Results []Item `json:"results"`
}

type Item struct {
	ID   string `json:"_id"`
	Data struct {
		MAP struct {
			Value string `json:"value"`
		} `json:"MAP"`
		Armor struct {
			Value string `json:"value,omitempty"`
		} `json:"armor,omitempty"`
		BaseItem string `json:"baseItem"`
		Bonus    struct {
			Value interface{} `json:"value"`
		} `json:"bonus"`
		BonusDamage struct {
			Value int `json:"value"`
		} `json:"bonusDamage"`
		BrokenThreshold struct {
			Value int `json:"value"`
		} `json:"brokenThreshold"`
		BulkCapacity struct {
			Value string `json:"value"`
		} `json:"bulkCapacity"`
		Category  string `json:"category"`
		Collapsed struct {
			Value bool `json:"value"`
		} `json:"collapsed"`
		ContainerID struct {
			Value string `json:"value"`
		} `json:"containerId"`
		Damage struct {
			DamageType string `json:"damageType"`
			Dice       int    `json:"dice"`
			Die        string `json:"die"`
			Value      string `json:"value"`
		} `json:"damage"`
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
		Equipped struct {
			Value bool `json:"value"`
		} `json:"equipped"`
		EquippedBulk struct {
			Value string `json:"value"`
		} `json:"equippedBulk"`
		Group interface{} `json:"group"`
		Hands struct {
			Value interface{} `json:"value"`
		} `json:"hands"`
		Hardness struct {
			Value int `json:"value"`
		} `json:"hardness"`
		Hp struct {
			Value int `json:"value"`
		} `json:"hp"`
		Identification struct {
			Status       string `json:"status"`
			Unidentified struct {
				Data struct {
					Description struct {
						Value string `json:"value"`
					} `json:"description"`
				} `json:"data"`
				Img  string `json:"img"`
				Name string `json:"name"`
			} `json:"unidentified"`
		} `json:"identification"`
		Invested struct {
			Value bool `json:"value"`
		} `json:"invested"`
		Level struct {
			Value int `json:"value"`
		} `json:"level"`
		MaxHp struct {
			Value int `json:"value"`
		} `json:"maxHp"`
		NegateBulk struct {
			Value string `json:"value"`
		} `json:"negateBulk"`
		PotencyRune struct {
			Value interface{} `json:"value"`
		} `json:"potencyRune"`
		PreciousMaterial struct {
			Value interface{} `json:"value"`
		} `json:"preciousMaterial"`
		PreciousMaterialGrade struct {
			Value interface{} `json:"value"`
		} `json:"preciousMaterialGrade"`
		Price struct {
			Value interface{} `json:"value"`
		} `json:"price"`
		Property1 struct {
			CritDamage             string      `json:"critDamage"`
			CritDamageType         string      `json:"critDamageType"`
			CritDice               interface{} `json:"critDice"`
			CritDie                string      `json:"critDie"`
			CriticalConditionType  string      `json:"criticalConditionType"`
			CriticalConditionValue interface{} `json:"criticalConditionValue"`
			DamageType             string      `json:"damageType"`
			Dice                   interface{} `json:"dice"`
			Die                    string      `json:"die"`
			StrikeConditionType    string      `json:"strikeConditionType"`
			StrikeConditionValue   interface{} `json:"strikeConditionValue"`
			Value                  string      `json:"value"`
		} `json:"property1"`
		PropertyRune1 struct {
			Value interface{} `json:"value"`
		} `json:"propertyRune1"`
		PropertyRune2 struct {
			Value interface{} `json:"value"`
		} `json:"propertyRune2"`
		PropertyRune3 struct {
			Value interface{} `json:"value"`
		} `json:"propertyRune3"`
		PropertyRune4 struct {
			Value interface{} `json:"value"`
		} `json:"propertyRune4"`
		Quantity struct {
			Value int `json:"value"`
		} `json:"quantity"`
		Range  interface{} `json:"range"`
		Reload struct {
			Value string `json:"value"`
		} `json:"reload"`
		Rules []interface{} `json:"rules"`
		Size  struct {
			Value string `json:"value"`
		} `json:"size"`
		Source struct {
			Value string `json:"value"`
		} `json:"source"`
		Specific struct {
			Value bool `json:"value"`
		} `json:"specific"`
		SplashDamage struct {
			Value interface{} `json:"value"`
		} `json:"splashDamage"`
		StackGroup struct {
			Value string `json:"value"`
		} `json:"stackGroup"`
		StrikingRune struct {
			Value string `json:"value"`
		} `json:"strikingRune"`
		Traits struct {
			Custom string `json:"custom"`
			Rarity struct {
				Value string `json:"value"`
			} `json:"rarity"`
			Value []string `json:"value"`
		} `json:"traits"`
		Usage struct {
			Value string `json:"value"`
		} `json:"usage"`
		Weight struct {
			Value interface{} `json:"value"`
		} `json:"weight"`
	} `json:"data,omitempty"`
	Effects []interface{} `json:"effects"`
	Name    string        `json:"name"`
	Type    string        `json:"type"`
}

func NewItem() *Item {
	return &Item{}
}

type ArmorResp struct {
	Count   int     `json:"count"`
	Results []Armor `json:"results"`
}

type Armor struct {
	ID   string `json:"_id"`
	Data struct {
		Armor struct {
			Value int `json:"value"`
		} `json:"armor"`
		BaseItem        string `json:"baseItem"`
		BrokenThreshold struct {
			Value interface{} `json:"value"`
		} `json:"brokenThreshold"`
		BulkCapacity struct {
			Value string `json:"value"`
		} `json:"bulkCapacity"`
		Category string `json:"category"`
		Check    struct {
			Value int `json:"value"`
		} `json:"check"`
		Collapsed struct {
			Value bool `json:"value"`
		} `json:"collapsed"`
		ContainerID struct {
			Value string `json:"value"`
		} `json:"containerId"`
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
		Dex struct {
			Value int `json:"value"`
		} `json:"dex"`
		Equipped struct {
			Value bool `json:"value"`
		} `json:"equipped"`
		EquippedBulk struct {
			Value string `json:"value"`
		} `json:"equippedBulk"`
		Group    string `json:"group"`
		Hardness struct {
			Value interface{} `json:"value"`
		} `json:"hardness"`
		Hp struct {
			Value interface{} `json:"value"`
		} `json:"hp"`
		Identification struct {
			Status       string `json:"status"`
			Unidentified struct {
				Data struct {
					Description struct {
						Value string `json:"value"`
					} `json:"description"`
				} `json:"data"`
				Img  string `json:"img"`
				Name string `json:"name"`
			} `json:"unidentified"`
		} `json:"identification"`
		Invested struct {
			Value bool `json:"value"`
		} `json:"invested"`
		Level struct {
			Value int `json:"value"`
		} `json:"level"`
		MaxHp struct {
			Value interface{} `json:"value"`
		} `json:"maxHp"`
		NegateBulk struct {
			Value string `json:"value"`
		} `json:"negateBulk"`
		PotencyRune struct {
			Value int `json:"value"`
		} `json:"potencyRune"`
		PreciousMaterial struct {
			Value string `json:"value"`
		} `json:"preciousMaterial"`
		PreciousMaterialGrade struct {
			Value string `json:"value"`
		} `json:"preciousMaterialGrade"`
		Price struct {
			Value string `json:"value"`
		} `json:"price"`
		PropertyRune1 struct {
			Value string `json:"value"`
		} `json:"propertyRune1"`
		PropertyRune2 struct {
			Value string `json:"value"`
		} `json:"propertyRune2"`
		PropertyRune3 struct {
			Value string `json:"value"`
		} `json:"propertyRune3"`
		PropertyRune4 struct {
			Value string `json:"value"`
		} `json:"propertyRune4"`
		Quantity struct {
			Value int `json:"value"`
		} `json:"quantity"`
		ResiliencyRune struct {
			Value string `json:"value"`
		} `json:"resiliencyRune"`
		Rules []interface{} `json:"rules"`
		Size  struct {
			Value string `json:"value"`
		} `json:"size"`
		Source struct {
			Value string `json:"value"`
		} `json:"source"`
		Speed struct {
			Value int `json:"value"`
		} `json:"speed"`
		StackGroup struct {
			Value string `json:"value"`
		} `json:"stackGroup"`
		Strength struct {
			Value int `json:"value"`
		} `json:"strength"`
		Traits struct {
			Custom string `json:"custom"`
			Rarity struct {
				Value string `json:"value"`
			} `json:"rarity"`
			Value []interface{} `json:"value"`
		} `json:"traits"`
		Usage struct {
			Value string `json:"value"`
		} `json:"usage"`
		Weight struct {
			Value string `json:"value"`
		} `json:"weight"`
	} `json:"data"`
	Effects []interface{} `json:"effects"`
	Name    string        `json:"name"`
	Type    string        `json:"type"`
}

func NewArmor() *Armor {
	return &Armor{}
}

type WeaponResp struct {
	Count   int      `json:"count"`
	Results []Weapon `json:"results"`
}

type Weapon struct {
	ID   string `json:"_id"`
	Data struct {
		MAP struct {
			Value string `json:"value"`
		} `json:"MAP"`
		BaseItem string `json:"baseItem"`
		Bonus    struct {
			Value interface{} `json:"value"`
		} `json:"bonus"`
		BonusDamage struct {
			Value int `json:"value"`
		} `json:"bonusDamage"`
		BrokenThreshold struct {
			Value int `json:"value"`
		} `json:"brokenThreshold"`
		BulkCapacity struct {
			Value string `json:"value"`
		} `json:"bulkCapacity"`
		Category  string `json:"category"`
		Collapsed struct {
			Value bool `json:"value"`
		} `json:"collapsed"`
		ContainerID struct {
			Value string `json:"value"`
		} `json:"containerId"`
		Damage struct {
			DamageType string `json:"damageType"`
			Dice       int    `json:"dice"`
			Die        string `json:"die"`
			Value      string `json:"value"`
		} `json:"damage"`
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
		Equipped struct {
			Value bool `json:"value"`
		} `json:"equipped"`
		EquippedBulk struct {
			Value string `json:"value"`
		} `json:"equippedBulk"`
		Group interface{} `json:"group"`
		Hands struct {
			Value interface{} `json:"value"`
		} `json:"hands"`
		Hardness struct {
			Value int `json:"value"`
		} `json:"hardness"`
		Hp struct {
			Value int `json:"value"`
		} `json:"hp"`
		Identification struct {
			Status       string `json:"status"`
			Unidentified struct {
				Data struct {
					Description struct {
						Value string `json:"value"`
					} `json:"description"`
				} `json:"data"`
				Img  string `json:"img"`
				Name string `json:"name"`
			} `json:"unidentified"`
		} `json:"identification"`
		Invested struct {
			Value bool `json:"value"`
		} `json:"invested"`
		Level struct {
			Value int `json:"value"`
		} `json:"level"`
		MaxHp struct {
			Value int `json:"value"`
		} `json:"maxHp"`
		NegateBulk struct {
			Value string `json:"value"`
		} `json:"negateBulk"`
		PotencyRune struct {
			Value interface{} `json:"value"`
		} `json:"potencyRune"`
		PreciousMaterial struct {
			Value string `json:"value"`
		} `json:"preciousMaterial"`
		PreciousMaterialGrade struct {
			Value string `json:"value"`
		} `json:"preciousMaterialGrade"`
		Price struct {
			Value interface{} `json:"value"`
		} `json:"price"`
		Property1 struct {
			CritDamage             string      `json:"critDamage"`
			CritDamageType         string      `json:"critDamageType"`
			CritDice               interface{} `json:"critDice"`
			CritDie                string      `json:"critDie"`
			CriticalConditionType  string      `json:"criticalConditionType"`
			CriticalConditionValue interface{} `json:"criticalConditionValue"`
			DamageType             string      `json:"damageType"`
			Dice                   interface{} `json:"dice"`
			Die                    string      `json:"die"`
			StrikeConditionType    string      `json:"strikeConditionType"`
			StrikeConditionValue   interface{} `json:"strikeConditionValue"`
			Value                  string      `json:"value"`
		} `json:"property1"`
		PropertyRune1 struct {
			Value string `json:"value"`
		} `json:"propertyRune1"`
		PropertyRune2 struct {
			Value string `json:"value"`
		} `json:"propertyRune2"`
		PropertyRune3 struct {
			Value string `json:"value"`
		} `json:"propertyRune3"`
		PropertyRune4 struct {
			Value string `json:"value"`
		} `json:"propertyRune4"`
		Quantity struct {
			Value int `json:"value"`
		} `json:"quantity"`
		Range  interface{} `json:"range"`
		Reload struct {
			Value string `json:"value"`
		} `json:"reload"`
		Rules []interface{} `json:"rules"`
		Size  struct {
			Value string `json:"value"`
		} `json:"size"`
		Source struct {
			Value string `json:"value"`
		} `json:"source"`
		SplashDamage struct {
			Value interface{} `json:"value"`
		} `json:"splashDamage"`
		StackGroup struct {
			Value string `json:"value"`
		} `json:"stackGroup"`
		StrikingRune struct {
			Value string `json:"value"`
		} `json:"strikingRune"`
		Traits struct {
			Custom string `json:"custom"`
			Rarity struct {
				Value string `json:"value"`
			} `json:"rarity"`
			Value []string `json:"value"`
		} `json:"traits"`
		Usage struct {
			Value string `json:"value"`
		} `json:"usage"`
		Weight struct {
			Value interface{} `json:"value"`
		} `json:"weight"`
	} `json:"data"`
	Effects []interface{} `json:"effects"`
	Name    string        `json:"name"`
	Type    string        `json:"type"`
}

func NewWeapon() *Weapon {
	return &Weapon{}
}

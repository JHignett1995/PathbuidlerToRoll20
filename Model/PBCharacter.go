package Model

type PBCharacter struct {
	Success bool `json:"success"`
	Build   struct {
		Name       string   `json:"name"`
		Class      string   `json:"class"`
		Level      int      `json:"level"`
		Ancestry   string   `json:"ancestry"`
		Heritage   string   `json:"heritage"`
		Background string   `json:"background"`
		Alignment  string   `json:"alignment"`
		Gender     string   `json:"gender"`
		Age        string   `json:"age"`
		Deity      string   `json:"deity"`
		Size       int      `json:"size"`
		Keyability string   `json:"keyability"`
		Languages  []string `json:"languages"`
		Attributes struct {
			Ancestryhp      int `json:"ancestryhp"`
			Classhp         int `json:"classhp"`
			Bonushp         int `json:"bonushp"`
			BonushpPerLevel int `json:"bonushpPerLevel"`
			Speed           int `json:"speed"`
			SpeedBonus      int `json:"speedBonus"`
		} `json:"attributes"`
		Abilities struct {
			Str int `json:"str"`
			Dex int `json:"dex"`
			Con int `json:"con"`
			Int int `json:"int"`
			Wis int `json:"wis"`
			Cha int `json:"cha"`
		} `json:"abilities"`
		Proficiencies struct {
			ClassDC       int `json:"classDC"`
			Perception    int `json:"perception"`
			Fortitude     int `json:"fortitude"`
			Reflex        int `json:"reflex"`
			Will          int `json:"will"`
			Heavy         int `json:"heavy"`
			Medium        int `json:"medium"`
			Light         int `json:"light"`
			Unarmored     int `json:"unarmored"`
			Advanced      int `json:"advanced"`
			Martial       int `json:"martial"`
			Simple        int `json:"simple"`
			Unarmed       int `json:"unarmed"`
			CastingArcane int `json:"castingArcane"`
			CastingDivine int `json:"castingDivine"`
			CastingOccult int `json:"castingOccult"`
			CastingPrimal int `json:"castingPrimal"`
			Acrobatics    int `json:"acrobatics"`
			Arcana        int `json:"arcana"`
			Athletics     int `json:"athletics"`
			Crafting      int `json:"crafting"`
			Deception     int `json:"deception"`
			Diplomacy     int `json:"diplomacy"`
			Intimidation  int `json:"intimidation"`
			Medicine      int `json:"medicine"`
			Nature        int `json:"nature"`
			Occultism     int `json:"occultism"`
			Performance   int `json:"performance"`
			Religion      int `json:"religion"`
			Society       int `json:"society"`
			Stealth       int `json:"stealth"`
			Survival      int `json:"survival"`
			Thievery      int `json:"thievery"`
		} `json:"proficiencies"`
		Feats                 [][]interface{} `json:"feats"`
		Specials              []string        `json:"specials"`
		Lores                 [][]interface{} `json:"lores"`
		Equipment             [][]interface{} `json:"equipment"`
		SpecificProficiencies struct {
			Trained   []interface{} `json:"trained"`
			Expert    []interface{} `json:"expert"`
			Master    []interface{} `json:"master"`
			Legendary []interface{} `json:"legendary"`
		} `json:"specificProficiencies"`
		Weapons []PbWeapon `json:"weapons"`
		Money   struct {
			Pp int `json:"pp"`
			Gp int `json:"gp"`
			Sp int `json:"sp"`
			Cp int `json:"cp"`
		} `json:"money"`
		Armor []struct {
			Name    string      `json:"name"`
			Qty     int         `json:"qty"`
			Prof    string      `json:"prof"`
			Pot     int         `json:"pot"`
			Res     interface{} `json:"res"`
			Mat     string      `json:"mat,omitempty"`
			Display string      `json:"display,omitempty"`
			Worn    bool        `json:"worn"`
			Runes   []string    `json:"runes"`
		} `json:"armor"`
		SpellCasters []Casting     `json:"spellCasters"`
		Formula      []interface{} `json:"formula"`
		Pets         []struct {
			Type            string          `json:"type"`
			Animal          string          `json:"animal"`
			Name            string          `json:"name"`
			Mature          bool            `json:"mature"`
			Incredible      bool            `json:"incredible"`
			IncredibleType  string          `json:"incredibleType"`
			Specializations []interface{}   `json:"specializations"`
			Armor           string          `json:"armor"`
			Equipment       [][]interface{} `json:"equipment"`
		} `json:"pets"`
		AcTotal struct {
			AcProfBonus    int    `json:"acProfBonus"`
			AcAbilityBonus int    `json:"acAbilityBonus"`
			AcItemBonus    int    `json:"acItemBonus"`
			AcTotal        int    `json:"acTotal"`
			ShieldBonus    string `json:"shieldBonus"`
		} `json:"acTotal"`
	} `json:"build"`
}

type PbWeapon struct {
	Name    string   `json:"name"`
	Qty     int      `json:"qty"`
	Prof    string   `json:"prof"`
	Die     string   `json:"die"`
	Pot     int      `json:"pot"`
	Str     string   `json:"str"`
	Mat     string   `json:"mat"`
	Display string   `json:"display"`
	Runes   []string `json:"runes"`
}

type Casting struct {
	Name             string `json:"name"`
	MagicTradition   string `json:"magicTradition"`
	SpellcastingType string `json:"spellcastingType"`
	Ability          string `json:"ability"`
	Proficiency      int    `json:"proficiency"`
	FocusPoints      int    `json:"focusPoints"`
	Spells           []struct {
		SpellLevel int      `json:"spellLevel"`
		List       []string `json:"list"`
	} `json:"spells"`
	PerDay []int `json:"perDay"`
}

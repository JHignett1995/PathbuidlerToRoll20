package Model

import (
	"PathbuidlerToRoll20/constants"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var counter = 0

func Convert(pb *PBCharacter) *Roll20Character {
	r20 := NewRoll20Character()
	convertCharacter(r20, pb)
	startingAttributes(r20, pb)
	languages(r20, pb)
	abilityScores(r20, pb)
	skills(r20, pb)
	weaponProficiency(r20, pb)
	armorProficiency(r20, pb)
	savingThrows(r20, pb)
	perception(r20, pb)
	classDC(r20, pb)
	classAC(r20, pb)
	characterHealth(r20, pb)
	allFeats(r20, pb)
	if len(pb.Build.SpellCasters) > 0 {
		spellSettings(r20, pb)
		spellsPerDay(r20, pb)
		spells(r20, pb)
	}
	sortItems(r20, pb)
	sortArmor(r20, pb)
	sortWeapons(r20, pb)
	sortLores(r20, pb)
	return r20
}

func convertCharacter(r20 *Roll20Character, pb *PBCharacter) {
	r20.Character.Name = pb.Build.Name
	r20.Type = "character"
	r20.SchemaVersion = 3
	r20.Character.Abilities = []string{}
}

func generateId() string {
	//$id = '-'.str_pad($counter, 5, '0', STR_PAD_LEFT).'-'.md5(uniqid('roll20', true));

	c := fmt.Sprintf("%05d", counter)
	hash := md5.Sum([]byte("roll20" + strconv.FormatInt(time.Now().UnixMicro(), 10)))
	counter++
	return "-" + c + "-" + hex.EncodeToString(hash[:])
}

func startingAttributes(r20 *Roll20Character, pb *PBCharacter) {
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("version_character", 4.5, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("name", pb.Build.Name, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("ancestry_heritage", pb.Build.Ancestry+" / "+pb.Build.Heritage, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("deity", pb.Build.Deity, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("class", pb.Build.Class, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("background", pb.Build.Background, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("alignment", pb.Build.Alignment, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("level", pb.Build.Level, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("age", pb.Build.Age, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("gender_pronouns", pb.Build.Gender, generateId()))
	var size = map[int]string{1: "Small", 2: "Medium", 3: "large"}
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("size", size[pb.Build.Size], generateId()))
}

func languages(r20 *Roll20Character, pb *PBCharacter) {

	for _, l := range pb.Build.Languages {
		langId := generateId()
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_languages_"+langId+"_language", l, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_languages_"+langId+"_toggles", "display, ", generateId()))
	}
}

func abilityScores(r20 *Roll20Character, pb *PBCharacter) {
	//str
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("strength", pb.Build.Abilities.Str, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("strength_score", strconv.Itoa(pb.Build.Abilities.Str), generateId()))
	mod, half := abilityMods(pb.Build.Abilities.Str)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("strength_modifier", mod, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("strength_modifier_half", half, generateId()))
	//dex
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("dexterity", pb.Build.Abilities.Dex, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("dexterity_score", strconv.Itoa(pb.Build.Abilities.Dex), generateId()))
	mod, half = abilityMods(pb.Build.Abilities.Dex)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("dexterity_modifier", mod, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("dexterity_modifier_half", half, generateId()))
	//con
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("constitution", pb.Build.Abilities.Con, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("constitution_score", strconv.Itoa(pb.Build.Abilities.Con), generateId()))
	mod, half = abilityMods(pb.Build.Abilities.Con)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("constitution_modifier", mod, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("constitution_modifier_half", half, generateId()))
	//int
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("intelligence", pb.Build.Abilities.Int, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("intelligence_score", strconv.Itoa(pb.Build.Abilities.Int), generateId()))
	mod, half = abilityMods(pb.Build.Abilities.Int)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("intelligence_modifier", mod, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("intelligence_modifier_half", half, generateId()))
	//wis
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("wisdom", pb.Build.Abilities.Wis, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("wisdom_score", strconv.Itoa(pb.Build.Abilities.Wis), generateId()))
	mod, half = abilityMods(pb.Build.Abilities.Wis)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("wisdom_modifier", mod, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("wisdom_modifier_half", half, generateId()))
	//cha
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("charisma", pb.Build.Abilities.Cha, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("charisma_score", strconv.Itoa(pb.Build.Abilities.Cha), generateId()))
	mod, half = abilityMods(pb.Build.Abilities.Cha)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("charisma_modifier", mod, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("charisma_modifier_half", half, generateId()))
}

func abilityMods(score int) (int, int) {
	var mod, half int
	a := score - 10

	if a%2 == 0 {
		mod = a / 2
	} else if a < 10 {
		a--
		mod = a / 2
	} else if a > 10 {
		a++
		mod = a / 2
	}
	half = mod / 2
	return mod, half
}

func skills(r20 *Roll20Character, pb *PBCharacter) {
	proficiencies := reflect.ValueOf(pb.Build.Proficiencies)

	// Let's get a slice of all skills
	keySlice := make([]string, 0)
	for key := range constants.Skills {
		keySlice = append(keySlice, key)
	}

	// Now sort the slice
	sort.Strings(keySlice)

	for _, skill := range keySlice {
		mod := constants.Skills[skill]
		profValue := 0
		if proficiencies.FieldByName(skill).CanInt() {
			profValue = int(proficiencies.FieldByName(skill).Int())
		}
		//profLabel := proficiencyLabel(profValue)
		abilities := reflect.ValueOf(pb.Build.Abilities)
		skillScore, profScore, abMod, bonus := calculateSkill(pb.Build.Level, profValue, int(abilities.FieldByName(mod).Int()), pb)

		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill, skillScore, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_proficiency_display", proficiencyLabel(profValue), generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_proficiency", profScore, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_item", bonus, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_temporary", "0", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_armor", "0", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_rank", profValue, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_ability", abMod, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute(skill+"_ability_select", constants.Abilities[mod], generateId()))
	}

}

func calculateSkill(rank, prof, abScore int, pb *PBCharacter) (int, int, int, int) {
	var skillScore, proficiencyScore, bonus = 0, 0, 0
	//Feat untrained improvisation U = 0.5*rank || rank when rank>=7
	mod, _ := abilityMods(abScore)

	if prof > 0 {
		proficiencyScore = prof + rank
	} else {
		//println(featFinder("Untrained Improvisation", pb))
		if featFinder("Untrained Improvisation", pb) {
			if rank >= 7 {
				bonus = rank
			} else {
				bonus = rank / 2
			}
		}
	}

	skillScore = mod + proficiencyScore + bonus
	return skillScore, proficiencyScore, mod, bonus
}

func featFinder(featName string, pb *PBCharacter) bool {
	for _, v := range pb.Build.Feats {
		if v[0] == featName {
			return true
		}
	}
	return false
}

func proficiencyLabel(v int) string {
	switch v {
	case 0:
		return "U"
	case 2:
		return "T"
	case 4:
		return "E"
	case 6:
		return "L"
	default:
		return "U"
	}
}

func weaponProficiency(r20 *Roll20Character, pb *PBCharacter) {
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("weapon_proficiencies_simple_rank", pb.Build.Proficiencies.Simple, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("weapon_proficiencies_martial_rank", pb.Build.Proficiencies.Martial, generateId()))
	id := generateId()
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_weapon-proficiencies_"+id+"_weapon_proficiencies_other", "Unarmed", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_weapon-proficiencies_"+id+"_weapon_proficiencies_other", pb.Build.Proficiencies.Unarmed, generateId()))
}

func armorProficiency(r20 *Roll20Character, pb *PBCharacter) {
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_unarmored_rank", pb.Build.Proficiencies.Unarmored, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_light_rank", pb.Build.Proficiencies.Light, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_medium_rank", pb.Build.Proficiencies.Medium, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_heavy_rank", pb.Build.Proficiencies.Heavy, generateId()))
}

func savingThrows(r20 *Roll20Character, pb *PBCharacter) {

	skillScore, profScore, _, _ := calculateSkill(pb.Build.Level, pb.Build.Proficiencies.Fortitude, pb.Build.Abilities.Con, pb)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_fortitude", skillScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_fortitude_proficiency", profScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_fortitude_proficiency_display", proficiencyLabel(pb.Build.Proficiencies.Fortitude), generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_fortitude_rank", pb.Build.Proficiencies.Fortitude, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_fortitude_ability_select", constants.Abilities["Con"], generateId()))

	skillScore, profScore, _, _ = calculateSkill(pb.Build.Level, pb.Build.Proficiencies.Reflex, pb.Build.Abilities.Dex, pb)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_reflex", skillScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_reflex_proficiency", profScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_reflex_proficiency_display", proficiencyLabel(pb.Build.Proficiencies.Reflex), generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_reflex_rank", pb.Build.Proficiencies.Reflex, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_reflex_ability_select", constants.Abilities["Dex"], generateId()))

	skillScore, profScore, _, _ = calculateSkill(pb.Build.Level, pb.Build.Proficiencies.Will, pb.Build.Abilities.Wis, pb)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_will", skillScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_will_proficiency", profScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_will_proficiency_display", proficiencyLabel(pb.Build.Proficiencies.Will), generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_will_rank", pb.Build.Proficiencies.Will, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("saving_throws_will_ability_select", constants.Abilities["Wis"], generateId()))

}

func perception(r20 *Roll20Character, pb *PBCharacter) {
	skillScore, profScore, _, _ := calculateSkill(pb.Build.Level, pb.Build.Proficiencies.Perception, 10, pb)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception", skillScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception_proficiency_display", proficiencyLabel(pb.Build.Proficiencies.Perception), generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception_proficiency", profScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception_item", 0, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception_temporary", "0", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception_rank", pb.Build.Proficiencies.Perception, generateId()))
}

func classDC(r20 *Roll20Character, pb *PBCharacter) {
	r, i := utf8.DecodeRuneInString(pb.Build.Keyability)
	pb.Build.Keyability = string(unicode.ToTitle(r)) + pb.Build.Keyability[i:]
	abs := reflect.ValueOf(pb.Build.Abilities)
	abScore := abs.FieldByName(pb.Build.Keyability).Int()

	skillScore, profScore, abMod, _ := calculateSkill(pb.Build.Level, pb.Build.Proficiencies.ClassDC, int(abScore), pb)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("class_dc", 10+skillScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("class_dc_rank", pb.Build.Proficiencies.ClassDC, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception_proficiency_display", proficiencyLabel(pb.Build.Proficiencies.ClassDC), generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("perception_proficiency", profScore, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("class_dc_key_ability", abMod, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("class_dc_key_ability_select", constants.Abilities[pb.Build.Keyability], generateId()))
}

func classAC(r20 *Roll20Character, pb *PBCharacter) {
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("ac", pb.Build.AcTotal.AcTotal, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class", pb.Build.AcTotal.AcTotal, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_ability", pb.Build.AcTotal.AcAbilityBonus, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_rank", pb.Build.AcTotal.AcProfBonus-pb.Build.Level, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_proficiency", pb.Build.AcTotal.AcProfBonus, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_proficiency_display", proficiencyLabel(pb.Build.AcTotal.AcProfBonus-pb.Build.Level), generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_ability_select", constants.Abilities["Dex"], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_dc_rank", strconv.Itoa(pb.Build.AcTotal.AcAbilityBonus), generateId()))
	shield, _ := strconv.Atoi(pb.Build.AcTotal.ShieldBonus)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("armor_class_shield", pb.Build.AcTotal.AcTotal+shield, generateId()))
}

func characterHealth(r20 *Roll20Character, pb *PBCharacter) {

	abMod, _ := abilityMods(pb.Build.Abilities.Con)
	hpMax := pb.Build.Attributes.Ancestryhp + ((pb.Build.Attributes.Classhp + abMod) * pb.Build.Level)
	//println(hpMax)
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("hit_points", hpMax, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("hit_points_max", hpMax, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("hit_points_class", pb.Build.Attributes.Classhp, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("hit_points_ancestry", pb.Build.Attributes.Ancestryhp, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("hit_points_notes", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("speed", pb.Build.Attributes.Speed+pb.Build.Attributes.SpeedBonus, generateId()))
}

type Feat struct {
	Name   string
	Desc   interface{}
	Source interface{}
	Level  interface{}
}

func isNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

func allFeats(r20 *Roll20Character, pb *PBCharacter) {

	var feats []Feat
	for _, f := range pb.Build.Feats {
		var feat Feat
		feat.Name = fmt.Sprint(f[0])
		feat.Desc = f[1]
		if len(f) > 2 {
			feat.Source = strings.ReplaceAll(strings.ToLower(fmt.Sprint(f[2])), " feat", "")
			if isNil(f[3]) {
				feat.Level = ""
			} else {
				feat.Level = f[3]
			}
		}
		feats = append(feats, feat)
	}
	for _, f := range feats {
		switch f.Source {
		case "class":
			fid := generateId()
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class", f.Name, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_type", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_level", f.Level, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_traits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_action", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_trigger", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_requirements", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_frequency", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_benefits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_feat_class_notes", constants.Feats[f.Name], generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-class_"+fid+"_toggles", "display, ", generateId()))
		case "ancestry":
			fid := generateId()
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry", f.Name, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_type", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_level", f.Level, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_traits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_action", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_trigger", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_requirements", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_frequency", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_benefits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_feat_ancestry_notes", constants.Feats[f.Name], generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-ancestry_"+fid+"_toggles", "display, ", generateId()))
		case "skill":
			fid := generateId()
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill", f.Name, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_type", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_level", f.Level, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_traits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_action", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_trigger", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_requirements", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_frequency", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_benefits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_feat_skill_notes", constants.Feats[f.Name], generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-skill_"+fid+"_toggles", "display, ", generateId()))
		case "general":
			fid := generateId()
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general", f.Name, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_type", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_level", f.Level, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_traits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_action", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_trigger", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_requirements", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_frequency", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_benefits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_feat_general_notes", constants.Feats[f.Name], generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-general_"+fid+"_toggles", "display, ", generateId()))
		default:
			fid := generateId()
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus", f.Name, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_type", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_level", f.Level, generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_traits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_action", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_trigger", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_requirements", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_frequency", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_benefits", "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_notes", constants.Feats[f.Name], generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_toggles", "display, ", generateId()))
		}
	}

	for _, f := range pb.Build.Specials {
		fid := generateId()
		feat, err := getFeat(f)
		println(f)
		if err != nil && strings.Contains(err.Error(), "Could Not Find Feat") {
			feat = NewFeat(f)
		}
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus", feat.Name, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_type", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_level", feat.Data.Level.Value, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_traits", feat.Data.Traits.Value, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_action", feat.Data.ActionType.Value, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_trigger", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_requirements", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_frequency", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_benefits", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_feat_bonus_notes", feat.Data.Description.Value, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("repeating_feat-bonus_"+fid+"_toggles", "display, ", generateId()))
	}
}

func getFeat(name string) (*APIFeat, error) {
	var feats FeatSearch
	cli := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.pathfinder2.fr/v1/pf2/spell", nil)
	constants.CheckErr(500, "New Spell Request: ", err)
	req.URL.Query().Set("name", name)
	req.Header.Set("Authorization", constants.PF2_API_KEY)

	resp, err := cli.Do(req)
	constants.CheckErr(500, "Performing Request: "+name, err)

	data, err := ioutil.ReadAll(resp.Body)
	constants.CheckErr(500, "Reading Body: ", err)

	err = json.Unmarshal(data, &feats)
	constants.CheckErr(500, "UnMarshaling Json: ", err)

	for _, feat := range feats.Results {
		if feat.Name == name {
			return &feat, nil
		}
	}
	return nil, errors.New("Could Not Find Feat")
}

func spellSettings(r20 *Roll20Character, pb *PBCharacter) {
	focusPoints := 0
	prepared := ""
	spontaneous := ""
	casting := make(map[string][]int)
	casting["divine"] = []int{0, 0}
	casting["arcane"] = []int{0, 0}
	casting["occult"] = []int{0, 0}
	casting["primal"] = []int{0, 0}

	for _, caster := range pb.Build.SpellCasters {

		if caster.FocusPoints > focusPoints {
			focusPoints = caster.FocusPoints
		}
		if caster.SpellcastingType == "prepared" {
			prepared = "prepared"
		}
		if caster.SpellcastingType == "spontaneous" {
			spontaneous = "spontaneous"
		}
		switch caster.MagicTradition {
		case "divine":
			casting[caster.MagicTradition] = []int{caster.Proficiency, caster.Proficiency + pb.Build.Level}
		case "arcane":
			casting[caster.MagicTradition] = []int{caster.Proficiency, caster.Proficiency + pb.Build.Level}
		case "occult":
			casting[caster.MagicTradition] = []int{caster.Proficiency, caster.Proficiency + pb.Build.Level}
		case "primal":
			casting[caster.MagicTradition] = []int{caster.Proficiency, caster.Proficiency + pb.Build.Level}
		}
	}

	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("spellcaster_prepared", prepared, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("spellcaster_spontaneous", spontaneous, generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_arcane_rank", casting["arcane"][0], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_arcane_proficiency", casting["arcane"][1], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_primal_rank", casting["primal"][0], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_primal_proficiency", casting["primal"][1], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_occult_rank", casting["occult"][0], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_occult_proficiency", casting["occult"][1], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_divine_rank", casting["divine"][0], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute("magic_tradition_divine_proficiency", casting["divine"][1], generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("focus_points", focusPoints, focusPoints, generateId()))

}

func spellsPerDay(r20 *Roll20Character, pb *PBCharacter) {
	levels := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	labels := []string{"cantrips", "level_1", "level_2", "level_3", "level_4", "level_5", "level_6", "level_7", "level_8", "level_9", "level_10"}
	for _, sc := range pb.Build.SpellCasters {
		for i, l := range levels {
			if sc.PerDay[i] > l {
				levels[i] = sc.PerDay[i]
			}
		}
	}

	for k, v := range levels {
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1(labels[k]+"_per_day", v, v, generateId()))
	}

}

func spellSetter(r20 *Roll20Character, inCasting []Spell, spellAbility string, spellAtk, spellDC int, label string) {

	for _, spell := range inCasting {
		id := generateId()
		var components []string

		fields := reflect.TypeOf(spell.Data.Components)
		values := reflect.ValueOf(spell.Data.Components)
		for i := 0; i < fields.NumField(); i++ {
			if values.Field(i).Bool() {
				components = append(components, fields.Name())
			}
		}

		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_name", spell.Name, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_spelllevel", spell.Data.Level, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_current_level", spell.Data.Level, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_traits", spell.Data.Traits.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_cast", components, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_target", spell.Data.Target.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_duration", spell.Data.Duration.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_range", spell.Data.Range.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_school", spell.Data.School.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_cast_actions", spell.Data.Time.Value+"-Action", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_description", strings.ReplaceAll(strings.ReplaceAll(spell.Data.Description.Value, "[[", ""), "]]", ""), "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_uses", 1, 1, generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_toggles", "display, ", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_spellattack", spellAtk, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_spelldc", spellDC, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_damage_critical_roll", constants.SpellCritRoll, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_damage_roll", constants.SpellDmgRoll, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_roll_critical_damage", constants.SpellCritDmg, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_magic_tradition", spell.Data.Traditions.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_attack_checkbox", "@{attack_roll}", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_damage_checkbox", "@{damage_roll}", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_damage_dice", spell.Data.Damage.Value.Num0.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_damage_type", spell.Data.Damage.Value.Num0.Type.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_"+label+"_"+id+"_damage_ability", spellAbility, "", generateId()))

	}
}

var spellClasses = []string{
	"Bard",
	"Champion",
	"Cleric",
	"Druid",
	"Magus",
	"Monk",
	"Oracle",
	"Sorcerer",
	"Summoner",
	"Witch",
	"Wizard",
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

func spells(r20 *Roll20Character, pb *PBCharacter) {
	var core []Casting
	var focus []Casting
	var innate []Casting

	for _, group := range pb.Build.SpellCasters {
		if contains(spellClasses, group.Name) {
			//println("Adding to Core " + group.Name)
			core = append(core, group)
		} else if group.Name == "Focus Spells" {
			//println("Adding to focus " + group.Name)
			focus = append(focus, group)
		} else {
			//println("Adding to innate " + group.Name)
			innate = append(innate, group)
		}
	}

	spell_AR_DCR := pb.Build.SpellCasters[0].Proficiency
	spellAbility := pb.Build.SpellCasters[0].Ability
	abs := reflect.ValueOf(pb.Build.Abilities)
	spellAbText := strings.Title(spellAbility)
	spellAbCode := constants.Abilities[spellAbText]

	skill, profScore, mod, _ := calculateSkill(pb.Build.Level, spell_AR_DCR, int(abs.FieldByName(spellAbText).Int()), pb)

	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_attack_rank", spell_AR_DCR, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_dc_rank", spell_AR_DCR, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_dc_key_ability", mod, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_attack_key_ability", mod, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_dc_key_ability_select", spellAbCode, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_attack_key_ability_select", spellAbCode, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("query_roll_critical_damage_dice", "0", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_attack_proficiency", profScore, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_attack_proficiency_display", proficiencyLabel(spell_AR_DCR), "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_attack", skill, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_dc_proficiency", profScore, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_dc_proficiency_display", proficiencyLabel(spell_AR_DCR), "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("spell_dc", 10+skill, "", generateId()))

	var inSpells, foSpells, coSpells, cantrips []Spell

	//fmt.Printf("%+v", core)
	//fmt.Printf("%+v", innate)
	//fmt.Printf("%+v", focus)
	for _, castSpells := range innate {
		for _, rank := range castSpells.Spells {
			if rank.SpellLevel == 0 {
				for _, spell := range rank.List {
					sp, err := getSpell(spell)
					if err != nil {
						panic(err)
					}
					cantrips = append(cantrips, *sp)
				}
			} else {
				for _, spell := range rank.List {
					sp, err := getSpell(spell)
					if err != nil {
						panic(err)
					}
					inSpells = append(inSpells, *sp)
				}
			}
		}
	}

	spellSetter(r20, inSpells, spellAbCode, skill, 10+skill, "spellinnate")

	for _, castSpells := range focus {
		for _, rank := range castSpells.Spells {
			if rank.SpellLevel == 0 {
				for _, spell := range rank.List {
					sp, err := getSpell(spell)
					if err != nil {
						panic(err)
					}
					cantrips = append(cantrips, *sp)
				}
			} else {
				for _, spell := range rank.List {
					sp, err := getSpell(spell)
					if err != nil {
						panic(err)
					}
					foSpells = append(foSpells, *sp)
				}
			}
		}
	}
	spellSetter(r20, foSpells, spellAbCode, skill, 10+skill, "spellfocus")

	for _, castSpells := range core {
		for _, rank := range castSpells.Spells {
			if rank.SpellLevel == 0 {
				for _, spell := range rank.List {
					sp, err := getSpell(spell)
					if err != nil && strings.Contains(err.Error(), "Could Not Find Spell") {
						sp = NewSpell(spell)
					}
					cantrips = append(cantrips, *sp)
				}
			} else {
				for _, spell := range rank.List {
					sp, err := getSpell(spell)
					if err != nil && strings.Contains(err.Error(), "Could Not Find Spell") {
						sp = NewSpell(spell)
					}
					coSpells = append(coSpells, *sp)
				}
			}
		}
	}

	// Populate Spellbook
	spellSetter(r20, coSpells, spellAbCode, skill, 10+skill, "normalspells")
	spellSetter(r20, cantrips, spellAbCode, skill, 10+skill, "cantrip")
}

func getSpell(name string) (*Spell, error) {
	constants.CheckErr(200, "Finding Spell: "+name, nil)
	var spells SpellSearch
	cli := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.pathfinder2.fr/v1/pf2/spell", nil)
	constants.CheckErr(500, "New Spell Request: ", err)
	req.URL.Query().Set("name", name)
	req.Header.Set("Authorization", constants.PF2_API_KEY)

	resp, err := cli.Do(req)
	constants.CheckErr(500, "Performing Request: "+name, err)

	data, err := ioutil.ReadAll(resp.Body)
	constants.CheckErr(500, "Reading Body: ", err)

	err = json.Unmarshal(data, &spells)
	constants.CheckErr(500, "UnMarshaling Json: ", err)

	for _, spell := range spells.Results {
		if spell.Name == name {
			return &spell, nil
		}
	}
	return nil, errors.New("Could Not Find Spell")
}

func sortItems(r20 *Roll20Character, pb *PBCharacter) {
	itemMap := make(map[string]float64)

	for _, item := range pb.Build.Equipment {
		itemMap[fmt.Sprint(item[0])] = item[1].(float64)
	}
	//fmt.Printf("%+v", itemMap)

	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("cp", pb.Build.Money.Pp, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("sp", pb.Build.Money.Cp, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("pp", pb.Build.Money.Sp, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("gp", pb.Build.Money.Gp, "", generateId()))

	for k, v := range itemMap {
		var item *Item
		item, err := getItem(k)
		if err != nil {
			if strings.Contains(err.Error(), "Find Item") {
				item = NewItem()
				item.Data.Description.Value = ""
				item.Data.Weight.Value = "0"
				item.Data.Price.Value = "0"
				item.Data.Level.Value = 0
			} else {
				constants.CheckErr(500, "Getting Item", err)
			}
		}
		id := generateId()
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-other_"+id+"_other_item", k, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-other_"+id+"_other_price", item.Data.Price.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-other_"+id+"_other_bulk", item.Data.Weight.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-other_"+id+"_other_misc", item.Data.Description.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-other_"+id+"_other_toggles", "display, ", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-other_"+id+"_other_quantity", v, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-other_"+id+"_other_level", item.Data.Level.Value, "", generateId()))
	}

}

func sortArmor(r20 *Roll20Character, pb *PBCharacter) {
	for _, armor := range pb.Build.Armor {
		id := generateId()
		item, err := getArmor(armor.Name)
		if err != nil && strings.Contains(err.Error(), "Could Not Find Item") {
			item = NewArmor()
		} else {
			constants.CheckErr(500, "Weapon Search Results: ", err)
		}
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-worn_"+id+"_worn_item", armor.Display, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-worn_"+id+"_worn_price", item.Data.Price.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-worn_"+id+"_worn_bulk", item.Data.Weight.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-worn_"+id+"_worn_level", item.Data.Level.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-worn_"+id+"_worn_misc", item.Data.Description, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-worn_"+id+"_worn_bulk_quantity", item.Data.Weight.Value, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_items-worn_"+id+"_toggles", "display, ", "", generateId()))
		if strings.Contains(strings.ToLower(armor.Name), "armor") {
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_armor_name", armor.Display, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_traits", item.Data.Traits.Rarity.Value, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_item", item.Data.Armor.Value, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_penalty_speed", item.Data.Speed.Value, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_strength", item.Data.Strength.Value, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_penalty_check", item.Data.Check.Value, "", generateId()))

		} else if strings.Contains(strings.ToLower(armor.Prof), "shield") {
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_shield_notes", item.Data.Description, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_shield_name", armor.Display, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_shield_ac_bonus", item.Data.Armor.Value, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_shield_temporary", 0, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_shield_hardness", item.Data.Hardness.Value, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_shield_bt", item.Data.BrokenThreshold.Value, "", generateId()))
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("armor_class_shield_hp", item.Data.Hp.Value, "", generateId()))
		}
	}
}

func getItem(name string) (*Item, error) {
	//constants.CheckErr(200, "Finding Item: "+name, nil)
	var items ItemResp
	cli := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.pathfinder2.fr/v1/pf2/equipment", nil)
	constants.CheckErr(500, "New Item Request: ", err)
	req.URL.Query().Set("name", name)
	req.Header.Set("Authorization", constants.PF2_API_KEY)

	resp, err := cli.Do(req)
	if err != nil && strings.Contains(err.Error(), "i/o timeout") {
		time.Sleep(time.Second * 1)
		resp, err = cli.Do(req)
	}
	constants.CheckErr(500, "Performing Request: "+name, err)

	data, err := ioutil.ReadAll(resp.Body)
	constants.CheckErr(500, "Reading Body: ", err)

	err = json.Unmarshal(data, &items)
	constants.CheckErr(300, "UnMarshaling Json: ", err)

	for _, item := range items.Results {
		if item.Name == name {
			return &item, nil
		}
	}
	return nil, errors.New("Could Not Find Item")
}

func getArmor(name string) (*Armor, error) {
	//constants.CheckErr(200, "Finding Item: "+name, nil)
	var items ArmorResp
	cli := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.pathfinder2.fr/v1/pf2/equipment", nil)
	constants.CheckErr(500, "New Item Request: ", err)
	req.URL.Query().Set("name", name)
	req.Header.Set("Authorization", constants.PF2_API_KEY)

	resp, err := cli.Do(req)
	constants.CheckErr(500, "Performing Request: "+name, err)

	data, err := ioutil.ReadAll(resp.Body)
	constants.CheckErr(500, "Reading Body: ", err)

	err = json.Unmarshal(data, &items)
	constants.CheckErr(300, "UnMarshaling Json: ", err)

	for _, item := range items.Results {
		if item.Name == name {
			return &item, nil
		}
	}
	return nil, errors.New("Could Not Find Item")
}

func getWeapon(name string) (*Weapon, error) {
	//constants.CheckErr(200, "Finding Item: "+name, nil)
	var items WeaponResp
	cli := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.pathfinder2.fr/v1/pf2/equipment", nil)
	constants.CheckErr(500, "New Item Request: ", err)
	req.URL.Query().Set("name", name)
	req.Header.Set("Authorization", constants.PF2_API_KEY)

	resp, err := cli.Do(req)
	constants.CheckErr(500, "Performing Request: "+name, err)

	data, err := ioutil.ReadAll(resp.Body)
	constants.CheckErr(500, "Reading Body: ", err)

	err = json.Unmarshal(data, &items)
	constants.CheckErr(300, "UnMarshaling Json: ", err)

	for _, item := range items.Results {
		if item.Name == name {
			return &item, nil
		}
	}
	return nil, errors.New("Could Not Find Item")
}

func sortWeapons(r20 *Roll20Character, pb *PBCharacter) {
	for _, weapon := range pb.Build.Weapons {
		id := generateId()
		var item *Weapon
		item, err := getWeapon(weapon.Name)
		if err != nil && strings.Contains(err.Error(), "Could Not Find Item") {
			item = NewWeapon()
		} else {
			constants.CheckErr(500, "Weapon Search Results: ", err)
		}
		if strings.ToLower(fmt.Sprintf("%v", item.Data.Group)) == "bow" || strings.ToLower(fmt.Sprintf("%v", item.Data.Group)) == "dart" || strings.ToLower(fmt.Sprintf("%v", item.Data.Group)) == "sling" || strings.ToLower(fmt.Sprintf("%v", item.Data.Group)) == "bomb" {
			rangedWeapons(r20, pb, id, item, weapon)
		} else {
			meleeWeapons(r20, pb, id, item, weapon)
		}

	}
}

func rangedWeapons(r20 *Roll20Character, pb *PBCharacter, id string, item *Weapon, weapon PbWeapon) {
	v := reflect.ValueOf(pb.Build.Proficiencies)
	var weaponProf int
	if !v.FieldByName(weapon.Prof).CanInt() {
		weaponProf = 0
	} else {
		weaponProf = int(v.FieldByName(weapon.Prof).Int())
	}

	str, _ := abilityMods(pb.Build.Abilities.Str)
	weaponScore, proficiencyScore, mod, _ := calculateSkill(pb.Build.Level, int(weaponProf), pb.Build.Abilities.Dex, pb)

	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon", weapon.Display, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_ability", mod, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_ability", str, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_proficiency", proficiencyScore, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_proficiency_display", proficiencyLabel(int(weaponProf)), "", generateId()))
	name := strings.ToLower(weapon.Display)
	if strings.Contains(name, "+") {
		atk, _ := strconv.Atoi(string(name[1]))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_strike", weaponScore+atk, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_item", atk, "", generateId()))
	} else {
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_strike", weaponScore, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_item", 0, "", generateId()))
	}
	var dmgDieCount int
	if strings.Contains(name, "striking") {
		if strings.Contains(name, "greater striking") {
			dmgDieCount = 3
		} else if strings.Contains(name, "major striking") {
			dmgDieCount = 4
		} else {
			dmgDieCount = 2
		}
	} else {
		dmgDieCount = 1
	}
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"damage_dice", dmgDieCount, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_dice_size", weapon.Die, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_display", weapon.Display, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_display", strconv.Itoa(dmgDieCount)+weapon.Die+strconv.Itoa(str)+item.Data.Damage.DamageType, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_bonus_display", item.Data.BonusDamage.Value, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_info", item.Data.Damage.DamageType, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_dice_query", "@{damage_dice}", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_roll", constants.RangedWeaponRoll1, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_roll2", constants.RangedWeaponRoll2, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_roll3", constants.RangedWeaponRoll3, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_roll", constants.RangedDamagedRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_critical_damage_dice", 2, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_critical_damage_dice_size", item.Data.Damage.Die, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_critical_damage_dice_query", "@{critical_damage_dice}", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_damage_critical_roll", constants.RangedCritDamageRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_roll_damage_roll", constants.RangedStrikeDamageRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_roll_critical_damage", constants.RangedStrikeCritRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_ability_select", "dex", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_rank", weaponProf, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_traits", item.Data.Traits.Value, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_range", item.Data.Range, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_weapon_reload", item.Data.Reload.Value, "", generateId()))
	for _, f := range item.Data.Traits.Value {
		if strings.Contains(strings.ToLower(f), "deadly") {
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_critical_damage_additional", strings.ReplaceAll(f, "deadly-", ""), "", generateId()))
			break
		}
	}
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_ranged-strikes_"+id+"_toggles", item.Data.Hp.Value, "", generateId()))
}

func meleeWeapons(r20 *Roll20Character, pb *PBCharacter, id string, item *Weapon, weapon PbWeapon) {
	v := reflect.ValueOf(pb.Build.Proficiencies)
	var weaponProf int
	if !v.FieldByName(weapon.Prof).CanInt() {
		weaponProf = 0
	} else {
		weaponProf = int(v.FieldByName(weapon.Prof).Int())
	}
	str, _ := abilityMods(pb.Build.Abilities.Str)
	weaponScore, proficiencyScore, mod, _ := calculateSkill(pb.Build.Level, weaponProf, pb.Build.Abilities.Dex, pb)

	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon", weapon.Display, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_ability", mod, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_ability", str, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_proficiency", proficiencyScore, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_proficiency_display", proficiencyLabel(weaponProf), "", generateId()))
	name := strings.ToLower(weapon.Display)
	if strings.Contains(name, "+") {
		atk, _ := strconv.Atoi(string(name[1]))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_strike", weaponScore+atk, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_item", atk, "", generateId()))
	} else {
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_strike", weaponScore, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_item", 0, "", generateId()))
	}
	var dmgDieCount int
	if strings.Contains(name, "striking") {
		if strings.Contains(name, "greater striking") {
			dmgDieCount = 3
		} else if strings.Contains(name, "major striking") {
			dmgDieCount = 4
		} else {
			dmgDieCount = 2
		}
	} else {
		dmgDieCount = 1
	}
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"damage_dice", dmgDieCount, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_dice_size", weapon.Die, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_display", weapon.Display, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_display", strconv.Itoa(dmgDieCount)+weapon.Die+strconv.Itoa(str)+item.Data.Damage.DamageType, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_bonus_display", item.Data.BonusDamage.Value, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_info", item.Data.Damage.DamageType, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_dice_query", "@{damage_dice}", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_roll", constants.RangedWeaponRoll1, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_roll2", constants.RangedWeaponRoll2, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_roll3", constants.RangedWeaponRoll3, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_roll", constants.RangedDamagedRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_critical_damage_dice", 2, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_critical_damage_dice_size", item.Data.Damage.Die, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_critical_damage_dice_query", "@{critical_damage_dice}", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_damage_critical_roll", constants.RangedCritDamageRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_roll_damage_roll", constants.RangedStrikeDamageRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_roll_critical_damage", constants.RangedStrikeCritRoll, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_ability_select", "dex", "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_rank", weaponProf, "", generateId()))
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_weapon_traits", item.Data.Traits.Value, "", generateId()))
	for _, f := range item.Data.Traits.Value {
		if strings.Contains(strings.ToLower(f), "deadly") {
			r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_critical_damage_additional", strings.ReplaceAll(f, "deadly-", ""), "", generateId()))
			break
		}
	}
	r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_melee-strikes_"+id+"_toggles", item.Data.Hp.Value, "", generateId()))
}

func sortLores(r20 *Roll20Character, pb *PBCharacter) {
	for _, lore := range pb.Build.Lores {
		id := generateId()
		skillScore, proficiencyScore, mod, _ := calculateSkill(pb.Build.Level, int(lore[1].(float64)), pb.Build.Abilities.Int, pb)
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_lore_name", lore[0], "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_lore_ability", mod, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_lore_proficiency", proficiencyScore, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_lore_proficiency_display", proficiencyLabel(int(lore[1].(float64))), "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_lore", skillScore, "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_lore_rank", lore[1], "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_lore_ability_select", "@{intelligence_modifier}", "", generateId()))
		r20.Character.Attribs = append(r20.Character.Attribs, *NewAttribute1("repeating_lore_"+id+"_toggles", "display, ", "", generateId()))
	}
}

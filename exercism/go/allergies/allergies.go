package allergies

var allAllergies = map[string]uint{}

func init() {
	allergies := []string{
		"eggs",
		"peanuts",
		"shellfish",
		"strawberries",
		"tomatoes",
		"chocolate",
		"pollen",
		"cats",
	}

	allAllergies = make(map[string]uint)
	for i, allergy := range allergies {
		allAllergies[allergy] = 1 << i
	}
}

func hasAllergy(allergyIndex uint, allergies uint) bool {
	return allergyIndex&allergies > 0
}
func Allergies(allergies uint) []string {
	result := make([]string, 0)
	for allergy, allergyIndex := range allAllergies {
		if hasAllergy(allergyIndex, allergies) {
			result = append(result, allergy)
		}
	}

	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	allergyIndex, ok := allAllergies[allergen]
	if !ok {
		panic("unknown allergy")
	}

	return hasAllergy(allergyIndex, allergies)
}

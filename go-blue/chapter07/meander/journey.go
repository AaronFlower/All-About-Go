package meander

import "strings"

type j struct {
	Name       string
	PlaceTypes []string
}

// Journeys defined the common journeys type.
var Journeys = []interface{}{
	j{Name: "Romantic", PlaceTypes: []string{
		"park",
		"bar",
		"movie_theater",
		"restaurant",
		"florist",
		"taxi_stand",
	}},
	j{
		Name: "Shopping",
		PlaceTypes: []string{
			"department_store",
			"cafe",
			"clothing_store",
			"jewelry_store",
			"shoe_store",
		},
	},
	j{Name: "Night Out", PlaceTypes: []string{
		"bar",
		"basino",
		"food",
		"night_club",
		"bar",
		"hospitla",
	}},
	j{
		Name: "Pamper",
		PlaceTypes: []string{
			"hair_care",
			"beauty_salon",
			"cafe",
			"spa",
		},
	},
}

func (j j) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}

package utils

import (
	"asuka/models"
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
func StrToType(s string) models.AnimeType {
	switch s {
	case "TV":
		return models.TV
	case "movie":
		return models.Movie
	case "OVA":
		return models.OVA
	case "Special":
		return models.Special
	case "ONA":
		return models.ONA
	case "Music":
		return models.Music
	}
	return models.InvalidType
}

func StrToSeason(s string) models.Season {
	switch s {
	case "fall":
		return models.Fall
	case "winter":
		return models.Winter
	case "summer":
		return models.Summer
	case "spring":
		return models.Spring
	}
	return models.InvalidSeason
}

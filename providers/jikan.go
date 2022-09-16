package providers

import (
	"asuka/models"
	"asuka/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JikanProvider struct {
	fmtUrl string
}

func NewJikanProvider() JikanProvider {
	return JikanProvider{fmtUrl: "https://api.jikan.moe/v4/anime?q=&quot;%s&quot;"}
}

func (j *JikanProvider) Search(query string, fuzzy bool) ([]models.Anime, error) {

	url := fmt.Sprintf(j.fmtUrl, query)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	animeResults := data["data"].([]interface{})

	var results []models.Anime

	for _, data := range animeResults {
		animeData := data.(map[string]interface{})
		var animeType models.AnimeType
		if animeData["type"] != nil {
			animeType = utils.StrToType(animeData["type"].(string))
		}
		// Exit early for "uncommon" entries
		if animeType == models.TV || animeType == models.Movie || animeType == models.Special {
			var animeId int
			if animeData["mal_id"] != nil {
				animeId = int(animeData["mal_id"].(float64))
			} else {
				continue
			}

			var posterURL string
			if animeData["images"] != nil {
				images := animeData["images"].(map[string]interface{})
				jpgImages := images["jpg"].(map[string]interface{})
				posterURL = jpgImages["image_url"].(string)
			}

			var titleEn string
			if animeData["title_english"] != nil {
				titleEn = animeData["title_english"].(string)
			}

			var title string
			if animeData["title"] != nil {
				title = animeData["title"].(string)
			} else {
				continue
			}

			var score float64
			if animeData["score"] != nil {
				score = animeData["score"].(float64)
			}

			var episodes int
			if animeData["episodes"] != nil {
				episodes = int(animeData["episodes"].(float64))
			} else {
				continue
			}

			var year int
			if animeData["year"] != nil {
				year = int(animeData["year"].(float64))
			}

			var season models.Season
			if animeData["season"] != nil {
				season = utils.StrToSeason(animeData["season"].(string))
			}

			anime := models.Anime{
				AnimeId:   animeId,
				PosterURL: posterURL,
				Type:      animeType,
				TitleEn:   titleEn,
				Title:     title,
				Score:     score,
				Episodes:  episodes,
				Year:      year,
				Season:    season,
			}

			results = append(results, anime)
		} else {
			continue
		}
	}
	return results, nil
}

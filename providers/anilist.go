package providers

import (
	"asuka/models"
	"asuka/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// API reference
const apiQueryFmtString = `{
Page(perPage:50){
pageInfo{
     total
     perPage
   }
   media(search: "%s", type: ANIME, sort: POPULARITY_DESC) {
     id
     title {
       romaji
       english
     }
	 format
     meanScore
	 episodes
     season
     seasonYear
     coverImage {
       large
     }
   }
 }
}`

const aniListUrl = "https://graphql.anilist.co"

type AniListProvider struct {
	queryFmt string
	baseUrl  string
}

type ReturnData struct {
	Data ApiResults `json:"data"`
}

type ApiResults struct {
	Page Page `json:"Page"`
}

type Page struct {
	PageInfo PageInfo       `json:"pageInfo"`
	Results  []AniListAnime `json:"media"`
}

type PageInfo struct {
	Total   float64 `json:"total"`
	PerPage float64 `json:"perPage"`
}

type AniListAnime struct {
	Id        float64      `json:"id"`
	PosterURL AniListImage `json:"coverImage"`
	Type      string       `json:"format"`
	Title     AniListTitle `json:"title"`
	Score     float64      `json:"meanScore"`
	Episodes  float64      `json:"episodes"`
	Year      float64      `json:"seasonYear"`
	Season    string       `json:"season"`
}

type AniListImage struct {
	LargeUrl string `json:"large"`
}

type AniListTitle struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}

func NewAniListProvider() AniListProvider {
	return AniListProvider{
		queryFmt: apiQueryFmtString,
		baseUrl:  aniListUrl,
	}
}

func (a AniListProvider) Search(query string, fuzzy bool) ([]models.Anime, error) {
	body, err := a.doRequest(query)
	if err != nil {
		return nil, err
	}

	var data ReturnData

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	var results []models.Anime

	for _, anime := range data.Data.Page.Results {
		results = append(results, anime.AniListToAnime())
	}

	if !fuzzy {
		var newResults []models.Anime

		for _, result := range results {
			if strings.Contains(strings.ToLower(result.Title), strings.ToLower(query)) {
				newResults = append(newResults, result)
			}
		}

		results = newResults
	}

	return results, nil
}

func (a AniListProvider) doRequest(query string) ([]byte, error) {
	requestBody, err := json.Marshal(map[string]string{
		"query": fmt.Sprintf(a.queryFmt, query),
	})

	if err != nil {
		return nil, err
	}

	timeout := 5 * time.Second

	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", a.baseUrl, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a AniListAnime) AniListToAnime() models.Anime {
	return models.Anime{
		AnimeId:   int(a.Id),
		PosterURL: a.PosterURL.LargeUrl,
		Type:      utils.StrToType(a.Type),
		TitleEn:   a.Title.English,
		Title:     a.Title.Romaji,
		Score:     a.Score / 10,
		Episodes:  int(a.Episodes),
		Year:      int(a.Year),
		Season:    utils.StrToSeason(strings.ToLower(a.Season)),
	}
}

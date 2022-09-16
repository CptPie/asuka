package models

type Season string

const (
	Winter        Season = "Winter"
	Fall                 = "Fall"
	Spring               = "Spring"
	Summer               = "Summer"
	InvalidSeason        = ""
)

type AnimeType string

const (
	TV          AnimeType = "TV"
	Movie                 = "Movie"
	OVA                   = "OVA"
	Special               = "Special"
	ONA                   = "ONA"
	Music                 = "Music"
	InvalidType           = ""
)

type Anime struct {
	AnimeId   int
	PosterURL string
	Type      AnimeType
	TitleEn   string
	Title     string
	Score     float64
	Episodes  int
	Year      int
	Season    Season
}

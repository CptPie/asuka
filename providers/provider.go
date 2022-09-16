package providers

import "asuka/models"

type Provider interface {
	Search(query string, fuzzy bool) ([]models.Anime, error)
}

package repositories

import (
	"net/http"

	"crud-backend/models"
)

const RANDOM_ARTICLE_URL = "https://en.wikipedia.org/wiki/Special:Random"

type WikipediaArticleRepository struct{}

func (r *WikipediaArticleRepository) GetRandom() (models.WikipediaArticle, error) {
	res, err := http.Get(RANDOM_ARTICLE_URL)

	item := models.WikipediaArticle{}

	if err != nil {
		return item, err
	}

	item.Url = res.Request.URL.String()

	return item, nil
}

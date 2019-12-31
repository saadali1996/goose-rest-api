package model

import (
	"github.com/advancedlogic/GoOse"
)

// Article description.
// swagger:model Article
type Article struct {

	// Title of the article
	//
	// required: true
	Title string `default:""`
	// URL of the article
	//
	// required: true
	Url string `default:""`
	// Top Image of the article
	//
	// required: true
	TopImage string `default:""`
}

type Error struct {
}

func GetArticleGoose(url string) Article {
	g := goose.New()

	article, _ := g.ExtractFromURL(url)

	article_obj := Article{
		Title:    "",
		Url:      "",
		TopImage: "",
	}

	if article != nil {

		article_obj.Url = url
		article_obj.Title = article.Title
		article_obj.TopImage = article.TopImage
	}
	return article_obj
}

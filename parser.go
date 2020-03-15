package nverify

import (
	"strings"

	"github.com/abadojack/whatlanggo"
	goose "github.com/advancedlogic/GoOse"
)

type Parser interface {
	ParseURL(URL string) (Article, error)
	DetectLanguage(content string) string
}

type GooseParser struct{}

func (parser GooseParser) ParseURL(URL string) (parsedArticle Article, err error) {
	g := goose.New()
	article, err := g.ExtractFromURL(URL)
	if err != nil {
		return
	}
	parsedArticle.Title = article.Title
	parsedArticle.Content = article.CleanedText
	parsedArticle.Description = article.MetaDescription
	parsedArticle.Keywords = strings.Split(article.MetaKeywords, ",")
	parsedArticle.Date = article.PublishDate
	parsedArticle.URL = URL
	return
}

func (parser GooseParser) DetectLanguage(content string) string {
	info := whatlanggo.Detect(content)
	return info.Lang.Iso6391()
}

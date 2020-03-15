package main

import (
	"log"
	"nverify"
	"os"
)

/*
	NVERIFY
	an API to verify a news for correlations and similar news from other sources

	-an URL is provided with get param "url"
	-an additional "accuracy" param can be provided, defines the amount of keywords extracted to use in the search
	-the content/keywords are extracted from the article using goose
	-search on NewsAPI for news with same keywords (limited to accuracy)
	-return a list of related news
*/
const (
	APIEnvKey = "NEWS_KEY"
)

func main() {
	apiKey := os.Getenv(APIEnvKey)
	if apiKey == "" {
		log.Fatalf("Need %v in env", APIEnvKey)
	}
	gparser := nverify.GooseParser{}
	newsSearch := nverify.NewsSearch{APIKey: apiKey}

	app := nverify.App{
		Parser: gparser,
		Search: newsSearch,
	}

	app.Start()
}

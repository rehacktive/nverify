package nverify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Search interface {
	Do(keywords []string, language string) ([]Article, error)
}

const (
	endpoint = "http://newsapi.org/v2/everything?q=%v&language=%v&apiKey=%v"
)

type NewsSearch struct {
	APIKey string
}

type searchResponse struct {
	Status       string
	TotalResults int
	Articles     []Article
}

func (api NewsSearch) Do(keywords []string, language string) (articles []Article, err error) {
	var client = &http.Client{
		Timeout: time.Second * 20,
	}
	finalURL := fmt.Sprintf(endpoint, encodeKeywords(keywords), language, api.APIKey)
	log.Println("search URL: ", finalURL)
	req, err := http.NewRequest(http.MethodGet, finalURL, nil)
	resp, err := client.Do(req)
	if err != nil || resp == nil {
		return
	}
	response := searchResponse{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		return
	}
	articles = response.Articles
	log.Println("found results #", len(articles))
	return
}

func encodeKeywords(keywords []string) string {
	merged := strings.Join(keywords, " AND ")
	return url.QueryEscape(merged)
}

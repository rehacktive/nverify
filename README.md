Search for correlated news

- create an account here: https://newsapi.org
- clone the project
- execute from command line:

`NEWS_KEY=your_api_key go run cmd/main.go`

- open the browser at http://localhost:8880

It's also an API, with curl for example:

`curl http://localhost:8880/verify?url=the_article_url`


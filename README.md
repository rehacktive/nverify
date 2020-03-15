**Search for correlated news**

The idea: every time I read a news and I'm not sure about the content, I search for sources (that I know) for the same content.
Then reading the results, I make my own idea about the fact.

This software is to help in this process: starting from an article (an URL), it will analyze it, extract keywords, and use them 
to search for similar articles from well-known sources.

**How it works**

- create an account here: https://newsapi.org
- clone the project
- execute from command line:

`NEWS_KEY=your_api_key go run cmd/main.go`

- open the browser at http://localhost:8880

It's also an API, with curl for example:

`curl http://localhost:8880/verify?url=the_article_url`


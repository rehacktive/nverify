package nverify

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	defaultAccuracy = 3
)

type App struct {
	Parser Parser
	Search Search

	server *http.Server
}

func (a App) Start() {
	r := mux.NewRouter()
	r.HandleFunc("/verify", a.verifyHandler)

	fs := http.FileServer(http.Dir("./web"))
	r.Handle("/", fs)

	a.server = &http.Server{Addr: ":8880", Handler: r}

	go a.server.ListenAndServe()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalf("unable to stop gracefully server: %v", err)
	}

	log.Println("best regards.")
}

func (a App) verifyHandler(w http.ResponseWriter, r *http.Request) {
	url, ok := r.URL.Query()["url"]

	if !ok || len(url[0]) == 0 {
		log.Println("'url' is missing")
		respondWithError(w, http.StatusNotAcceptable, "no url param provided")
		return
	}
	log.Println("Parsing: ", url[0])

	accuracy := defaultAccuracy
	accuracyParam, ok := r.URL.Query()["accuracy"]
	if ok && len(accuracyParam) > 0 {
		acc, err := strconv.Atoi(accuracyParam[0])
		if err != nil {
			log.Println("cannot parse accuracy: using default ", defaultAccuracy)
		} else {
			accuracy = acc
		}
	}

	article, err := a.Parser.ParseURL(url[0])
	if err != nil {
		log.Println("cannot parse the URL: ", err)
		respondWithError(w, http.StatusInternalServerError, "cannot parse the URL")
		return
	}
	log.Println("keywords found #", len(article.Keywords))
	if len(article.Keywords) == 0 || len(article.Keywords[0]) == 0 {
		respondWithError(w, http.StatusNotFound, "no keywords found")
		return
	}
	results, err := a.Search.Do(article.Keywords[:min(len(article.Keywords), accuracy)], a.Parser.DetectLanguage(article.Content))
	if err != nil {
		log.Println("cannot search news: ", err)
		respondWithError(w, http.StatusInternalServerError, "cannot search related")
		return
	}
	response := response{
		Total:    len(results),
		Articles: results,
	}
	respondWithJSON(w, http.StatusOK, response)
}

type response struct {
	Total    int
	Articles []Article
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload) // TODO convert to Encoder
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Println("error sending the response: ", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

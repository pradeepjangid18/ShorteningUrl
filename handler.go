package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func generateAndGetShortUrl(longUrl string, clientId string) string {
	shortUrl, isPresent := urlManager.getShortUrlIfAlreadyPresent(clientId, longUrl)
	if isPresent {
		return shortUrl
	}

	shortUrl = urlManager.generateShortUrl(clientId, longUrl)

	stats.updateUrlStats(longUrl)
	return shortUrl
}

func getLongUrlFromShortUrl(shortUrl string) (string, bool) {
	stats.updateUrlStats(shortUrl)
	return urlManager.getLongUrlIfAlreadyPresent(shortUrl)
}


func getShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	longUrl := vars["longUrl"]
	clientId := vars["clientId"]

	shortUrl := generateAndGetShortUrl(longUrl, clientId)
	w.WriteHeader(200)
	fmt.Fprintf(w, "shortUrl: " + shortUrl)
}

func getLongUrlHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	longUrl, ok := getLongUrlFromShortUrl(shortUrl)
	if !ok {
		w.WriteHeader(404)
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, "longUrl: " + longUrl)
	}
}
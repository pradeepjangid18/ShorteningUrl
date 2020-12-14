package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	urlManager = &urlData{ClientIdLongUrlToShortUrlMap: make(map[string]map[string]string), ShortUrlToLongUrlMap: make(map[string]string)}
	counter = &globalCounter{}
	stats = &urlStats{urlToCount: make(map[string]int)}
}

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/getShortenedUrl/{longUrl}/{clientId}", getShortUrlHandler)
	myRouter.HandleFunc("/getLongUrl/{shortUrl}", getLongUrlHandler)

	http.ListenAndServe(":8082", myRouter)
}
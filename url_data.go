package main

import "sync"

type urlData struct {
	mu sync.Mutex
	ClientIdLongUrlToShortUrlMap map[string]map[string]string
	ShortUrlToLongUrlMap map[string]string
}

var urlManager *urlData

func(u *urlData) getShortUrlIfAlreadyPresent(clientId string, longUrl string) (string, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, ok := u.ClientIdLongUrlToShortUrlMap[clientId]; !ok {
		return "", false
	}

	if _, ok := u.ClientIdLongUrlToShortUrlMap[clientId][longUrl]; !ok {
		return "", false
	}

	return u.ClientIdLongUrlToShortUrlMap[clientId][longUrl], true
}

func(u * urlData) getLongUrlIfAlreadyPresent(shortUrl string) (string, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	longUrl, ok := u.ShortUrlToLongUrlMap[shortUrl]
	if !ok {
		return "", false
	}
	return longUrl, true
}

func(u *urlData) generateShortUrl(clientId string, longUrl string) string {
	u.mu.Lock()
	defer u.mu.Unlock()

	_, ok := u.ClientIdLongUrlToShortUrlMap[clientId]
	if !ok {
		u.ClientIdLongUrlToShortUrlMap[clientId] = make(map[string]string)
	}

	counter := counter.incrementAndGet()
	shortUrl := idToShortUrl(counter)
	u.ClientIdLongUrlToShortUrlMap[clientId][longUrl] = shortUrl
	u.ShortUrlToLongUrlMap[shortUrl] = longUrl

	return shortUrl
}
package main

func idToShortUrl(id int64) string {
	var character string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var shortUrl string
	for {
		if id == 0 {
			break
		}
		ele := character[id%62]
		id = id/62
		shortUrl = string(ele) + shortUrl
	}
	return shortUrl
}
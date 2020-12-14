package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShortUrl", func() {


	//BeforeEach(func() {
	//	urlManager = &urlData{ClientIdLongUrlToShortUrlMap: make(map[string]map[string]string), ShortUrlToLongUrlMap: make(map[string]string)}
	//	counter = &globalCounter{}
	//	stats = &urlStats{urlToCount: make(map[string]int)}
	//})

	Describe("Test generateAndGetShortUrl", func() {
		Context("generate short url for long url : abcd and client id 1", func() {
			It("should return short url", func() {
				shortUrl := generateAndGetShortUrl("abcd", "1")
				Expect(shortUrl).To(Equal("b"))
			})
		})

		Context("generate again short url for long url : abcd and client id 1", func() {
			It("should return same short url", func() {
				shortUrl := generateAndGetShortUrl("abcd", "1")
				Expect(shortUrl).To(Equal("b"))
			})
		})

		Context("generate again short url for long url : abcd but client id 2", func() {
			It("should return same short url", func() {
				shortUrl := generateAndGetShortUrl("abcd", "2")
				Expect(shortUrl).To(Equal("c"))
			})
		})
	})

	Describe("Test getLongUrlFromShortUrl", func() {
		Context("generate long url for short url : zxyz", func() {
			It("should not return short url", func() {
				longUrl, isPresent := getLongUrlFromShortUrl("abcd")
				Expect(longUrl).To(Equal(""))
				Expect(isPresent).To(Equal(false))
			})
		})

		Context("generate long url for short url : b", func() {
			It("should not return short url", func() {
				longUrl, isPresent := getLongUrlFromShortUrl("b")
				Expect(longUrl).To(Equal("abcd"))
				Expect(isPresent).To(Equal(true))
			})
		})
	})

	Describe("Test urlStats", func() {

		Context("get count of url hit abcd", func() {
			It("should  return hit count 4", func() {
				count := stats.getUrlHitCount("abcd")
				Expect(count).To(Equal(3))
			})
		})

		Context("get count of url hit b", func() {
			It("should  return hit count 1", func() {
				count := stats.getUrlHitCount("b")
				Expect(count).To(Equal(1))
			})
		})
	})
})

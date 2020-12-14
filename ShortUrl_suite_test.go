package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestShortUrl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ShortUrl Suite")
}

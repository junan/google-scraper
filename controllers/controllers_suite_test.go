package controllers_test

import (
	"testing"
	_ "google-scraper/routers"
	_ "google-scraper/initializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

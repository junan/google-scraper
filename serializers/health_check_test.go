package serializers_test

import (
	"google-scraper/serializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheckSerializer", func() {
	Describe("#Data", func() {
		It("returns correct data", func() {
			serializer := serializers.HealthCheck{
				HealthCheck: true,
			}

			data := serializer.Data()

			Expect(data.Id).To(Equal(0))
			Expect(data.Success).To(Equal(true))
		})
	})
})

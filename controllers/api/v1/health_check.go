package apiv1

import "google-scraper/serializers"


type HealthCheck struct {
	baseAPIController
}

func (c *HealthCheck) Get() {
	healthCheckSerializer := serializers.HealthCheck{
		HealthCheck: true,
	}

	c.serveJSON(healthCheckSerializer.Data())
}

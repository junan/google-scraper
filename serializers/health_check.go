package serializers

type HealthCheck struct {
	HealthCheck bool
}

type healthCheckResponse struct {
	Id      int  `jsonapi:"primary,health_check"`
	Success bool `jsonapi:"attr,success"`
}

func (serializer *HealthCheck) Data() *healthCheckResponse {
	data := &healthCheckResponse{
		Success: serializer.HealthCheck,
	}

	return data
}

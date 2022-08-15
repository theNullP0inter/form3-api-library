package common

import "time"

const (
	TestEndPoint = "http://accountapi:8080"
	LiveEndPoint = "https://api.form3.tech"
)

type Config struct {
	Live        bool
	HttpTimeout time.Duration
}

func (c *Config) GetBaseUrl() string {
	if c.Live {
		return LiveEndPoint
	}
	return TestEndPoint
}

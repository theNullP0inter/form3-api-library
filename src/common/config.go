package common

const (
	TestEndPoint = "http://accountapi:8080"
	LiveEndPoint = "https://api.form3.tech"
)

type Config struct {
	Live bool
}

func (c *Config) GetBaseUrl() string {
	if c.Live {
		return LiveEndPoint
	}
	return TestEndPoint
}

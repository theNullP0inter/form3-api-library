package common

import "time"

var MockConfigTest = &Config{
	Live:        false,
	HttpTimeout: 10 * time.Second,
}

var MockConfigLive = &Config{
	Live:        true,
	HttpTimeout: 10 * time.Second,
}

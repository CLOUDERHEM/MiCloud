package request

import "time"

type Config struct {
	Timeout time.Duration
}

type Opt func(c *Config)

func Timeout(timeout time.Duration) Opt {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

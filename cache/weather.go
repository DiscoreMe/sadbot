package cache

import (
	"strings"
	"time"
)

const expirationWeather = 1 * time.Hour

func fixLocation(l string) string {
	return strings.ToLower(strings.ReplaceAll(l, " ", "_"))
}

func (c *Cache) WeatherGet(location string) ([]byte, error) {
	r, err := c.get(redisKeyWeather + fixLocation(location))
	if err != nil {
		return nil, err
	}
	return []byte(r), nil
}

func (c *Cache) WeatherSet(location string, data []byte) error {
	return c.r.Set(redisKeyWeather+fixLocation(location), string(data), expirationWeather).Err()
}

package pokeapi

import (
	"net/http"
	"time"

	"github.com/marcossnikel/pokecli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// Initialize a new http client with a specific timeout and cache
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

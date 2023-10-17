package wafris_traefik

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	URL string `json:"url,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		URL: "redis://localhost:6379?password=hello&protocol=3",
	}
}

// the main plugin struct.
type WafrisTraefikPlugin struct {
	next http.Handler
	url  string
	name string
}

// New created a new WafrisTraefikPlugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.URL) == 0 {
		return nil, fmt.Errorf("url cannot be empty")
	}

	return &WafrisTraefikPlugin{
		url:  config.URL,
		next: next,
		name: name,
	}, nil
}

func (a *WafrisTraefikPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// we only set a header for debugging purposes
	req.Header.Set("X-WafrisPlugin", "Passthru")

	// passthru everything for now
	a.next.ServeHTTP(rw, req)
}

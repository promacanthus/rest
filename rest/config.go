package rest

import (
	"fmt"
	"time"
	"os"

	"github.com/promacanthus/buildInfo/info"
)

type Config struct {
	// UserAgent is an optional field that specifies the caller of this request.
	UserAgent string `json:"user_agent,omitempty"`

	// Set retry parameters for http client.
	RetryCount       int           `json:"retry_count,omitempty"`
	RetryWaitTime    time.Duration `json:"retry_wait_time,omitempty"`
	RetryMaxWaitTime time.Duration `json:"retry_max_wait_time,omitempty"`

	BaseURLs map[string]string `json:"base_urls,omitempty"`

	// DataStore for calico internal client.
	DataStore string `json:"data_store,omitempty"`
}

// DefaultNephilaUserAgent returns the default user agent for Nephila.
func DefaultUserAgent() string {
	return buildUserAgent(
		info.Infos().GitCommit,
		info.Infos().GoOS,
		info.Infos().GoArch,
	)
}

// buildUserAgent is a Go function that constructs a user agent string.
// It takes the command, operating system, architecture, and commit as parameters and returns a string.
func buildUserAgent(commit, os, arch string) string {
	return fmt.Sprintf("%s/%s (%s/%s)", os.Args[0], commit, os, arch)
}

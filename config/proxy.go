package config

import (
	_ "embed"
	"fmt"
	"sync/atomic"

	"github.com/DOME-Marketplace/tmfgo/internal/errl"
	"github.com/goccy/go-yaml"
)

type UpstreamEntry struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Path string `yaml:"path"`
}

type UpstreamEntries map[string]UpstreamEntry

//go:embed proxy.yaml
var internalUpstreamYAMLContent []byte

type ProxyConfig struct {
	internalUpstreamEntries UpstreamEntries
}

func (p *ProxyConfig) GetUpstreamEntries() UpstreamEntries {
	return p.internalUpstreamEntries
}

var proxyConfig atomic.Pointer[ProxyConfig]

// GetProxyConfig returns the current proxy config, loading it from the embedded yaml file if not already loaded.
func GetProxyConfig() *ProxyConfig {
	cfg := proxyConfig.Load()
	// 	Parse the proxy yaml content only once at first time
	if cfg == nil {
		var newUpstreamEntries UpstreamEntries
		err := yaml.Unmarshal(internalUpstreamYAMLContent, &newUpstreamEntries)
		if err != nil {
			panic("error parsing proxy yaml content: " + err.Error())
		}
		cfg = UpdateProxyConfig(newUpstreamEntries)
	}
	return cfg
}

// UpdateProxyConfig updates in an atomic way the proxy config with new upstream entries.
func UpdateProxyConfig(newUpstreamEntries UpstreamEntries) *ProxyConfig {
	cfg := &ProxyConfig{
		internalUpstreamEntries: newUpstreamEntries,
	}
	proxyConfig.Store(cfg)
	return cfg
}

// InternalUpstreamURL returns a url of the form: http://hostname:port/path
func InternalUpstreamURL(resourceName string) (string, error) {

	cfg := GetProxyConfig()

	// Get the entry for the resource name
	entry, ok := cfg.internalUpstreamEntries[resourceName]
	if !ok {
		return "", errl.Errorf("unknown resource type: %s", resourceName)
	}

	// Build the URL, which is like: http://hostname:port/path
	url := fmt.Sprintf("http://%s:%d%s", entry.Host, entry.Port, entry.Path)
	return url, nil
}

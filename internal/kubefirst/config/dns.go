package config

import (
	"fmt"
)

type DNS struct {
	Provider                 DNSProvider       `yaml:"provider"`
	DomainFilters            []string          `yaml:"domainFilters"`
	Auth                     string            `yaml:"Auth"`
	Annotations              map[string]string `yaml:"annotations"`
	ExternalDNSHelmChartInfo ChartInfo         `yaml:"externalDNSHelmChartInfo"`
}

// DNSProvider represents the type of git provider
type DNSProvider string

const (
	DNSProviderCloudflare   DNSProvider = "cloudflare"
	DNSProviderAkamai       DNSProvider = "akamai"
	DNSProviderAWS          DNSProvider = "aws"
	DNSProviderAzure        DNSProvider = "azure"
	DNSProviderCivo         DNSProvider = "civo"
	DNSProviderDigitalOcean DNSProvider = "digitalocean"
	DNSProviderGoogle       DNSProvider = "google"
	DNSProviderVultr        DNSProvider = "vultr"
)

// DetermineDNSProvider figures out which provider is configured
func DetermineDNSProvider(dnsProvider DNSProvider) (DNSProvider, error) {

	switch DNSProvider(dnsProvider) {
	case DNSProviderCloudflare, DNSProviderAkamai, DNSProviderAWS, DNSProviderAzure, DNSProviderCivo, DNSProviderDigitalOcean, DNSProviderGoogle, DNSProviderVultr:
		return DNSProvider(dnsProvider), nil
	default:
		return "", fmt.Errorf("unsupported DNS provider: %s", dnsProvider)
	}
}

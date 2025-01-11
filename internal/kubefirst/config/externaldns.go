package config

type ExternalDNS struct {
	Provider                 string            `yaml:"provider"`
	DomainFilters            []string          `yaml:"domainFilters"`
	EnvName                  string            `yaml:"envName"`
	Annotations              map[string]string `yaml:"annotations"`
	ExternalDNSHelmChartInfo ChartInfo         `yaml:"externalDNSHelmChartInfo"`
}

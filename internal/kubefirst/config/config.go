package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	CloudProvider string       `yaml:"cloudProvider"`
	ClusterName   string       `yaml:"clusterName"`
	DomainName    string       `yaml:"domainName"`
	DNS           DNS          `yaml:"dns"`
	Git           Git          `yaml:"git"`
	GitopsConfig  GitopsConfig `yaml:"gitopsConfig"`
}

type GitopsConfig struct {
	RepoURL string
}

func ReadPlatformConfig() (*Config, error) {
	kubefirstConfig := "kubefirst.yml"
	platformConfig, err := os.ReadFile(kubefirstConfig)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	var config Config

	err = yaml.Unmarshal(platformConfig, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing kubefirst %q: %w", kubefirstConfig, err)
	}

	//! pull into isolate function
	gitprov, err := DetermineGitProvider(config.Git)
	if err != nil {
		return nil, fmt.Errorf("error determining git provider: %w", err)
	}

	switch gitprov {
	case ProviderGitHub:
		if config.Git.Auth == "https" {
			config.GitopsConfig.RepoURL = fmt.Sprintf("https://%v/%v/%v.git", "github.com", config.Git.GitHub.Organization, config.Git.GitHub.Repo)
		} else {
			config.GitopsConfig.RepoURL = fmt.Sprintf("git@%v:%v/%v.git", "github.com", config.Git.GitHub.Organization, config.Git.GitHub.Repo)
		}
	case ProviderGitLab:
		if config.Git.Auth == "https" {
			config.GitopsConfig.RepoURL = fmt.Sprintf("https://%v/%v/%v.git", "gitlab.com", config.Git.GitLab.Group, config.Git.GitLab.Repo)
		} else {
			config.GitopsConfig.RepoURL = fmt.Sprintf("git@%v:%v/%v.git", "gitlab.com", config.Git.GitLab.Group, config.Git.GitLab.Repo)
		}
	}

	if config.DNS.Provider == "" {
		provider, err := DetermineDNSProvider(DNSProvider(config.CloudProvider))
		if err != nil {
			return nil, fmt.Errorf("error determining DNS provider: %w", err)
		}
		config.DNS.Provider = provider
	}

	fmt.Printf("Git provider: %v\n", gitprov)
	fmt.Printf("DNS provider: %v\n", config.DNS.Provider)

	return &config, nil
}

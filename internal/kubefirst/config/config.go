package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	CloudProvider string       `yaml:"cloudProvider"`
	Cloud         Cloud        `yaml:"cloud"`
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

	err = hydrateCloudConfig(&config)
	if err != nil {
		return nil, fmt.Errorf("error hydrating cloud config: %w", err)
	}
	fmt.Printf("%+v", config)

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

func hydrateCloudConfig(cfg *Config) error {
	switch cfg.CloudProvider {
	case "aws":
		// TODO look these values up and hydrate the cloud object
		// get region
		region := "us-west-2"
		// get account id
		accountID := "126827061464"
		awsConfig := []byte(fmt.Sprintf("aws:\n  region: %v\n  accountID: %v\n", region, accountID))

		err := yaml.Unmarshal(awsConfig, &cfg.Cloud)
		if err != nil {
			return fmt.Errorf("error parsing aws config: %w", err)
		}
	// TODO constants CloudProvider
	case "azure":
		// TODO look these values up and hydrate the cloud object
		// get region
		region := "us-west-2"
		// get account id
		identityClientID := "MYAZURECLIENTID"
		azureConfig := []byte(fmt.Sprintf("azure:\n  region: %v\n  identityClientID: %v\n", region, identityClientID))

		err := yaml.Unmarshal(azureConfig, &cfg.Cloud)
		if err != nil {
			return fmt.Errorf("error parsing azure config: %w", err)
		}
	case "google":
		// TODO look these values up and hydrate the cloud object
		// get region
		region := "us-central1"
		// get account id
		projectName := "gcp-je"
		googleConfig := []byte(fmt.Sprintf("google:\n  region: %v\n  projectName: %v\n", region, projectName))

		err := yaml.Unmarshal(googleConfig, &cfg.Cloud)
		if err != nil {
			return fmt.Errorf("error parsing google config: %w", err)
		}
	}

	return nil
}

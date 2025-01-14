package config

import (
	"errors"
)

type Git struct {
	Auth   string `yaml:"auth"`
	GitHub GitHub `yaml:"github"`
	GitLab GitLab `yaml:"gitlab"`
}

type GitHub struct {
	Username     string `yaml:"username"`
	Organization string `yaml:"organization"`
	Repo         string `yaml:"repo"`
}

type GitLab struct {
	Username string `yaml:"username"`
	Group    string `yaml:"group"`
	Repo     string `yaml:"repo"`
}

// GitProvider represents the type of git provider
type GitProvider string

const (
	ProviderGitHub GitProvider = "github"
	ProviderGitLab GitProvider = "gitlab"
)

// DetermineGitProvider figures out which provider is configured
func DetermineGitProvider(git Git) (GitProvider, error) {
	// Check if GitHub is configured
	if git.GitHub != (GitHub{}) {
		return ProviderGitHub, nil
	}
	// Check if GitLab is configured
	if git.GitLab != (GitLab{}) {
		return ProviderGitLab, nil
	}
	return "", errors.New("no git provider configured")
}

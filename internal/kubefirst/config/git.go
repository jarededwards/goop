package config

import (
	"errors"
	"fmt"
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

// Example of using a switch statement with the provider
func HandleGitProvider(git Git) error {
	provider, err := DetermineGitProvider(git)
	if err != nil {
		return err
	}

	switch provider {
	case ProviderGitHub:
		return handleGitHub(git.GitHub)
	case ProviderGitLab:
		return handleGitLab(git.GitLab)
	default:
		return fmt.Errorf("unsupported git provider: %s", provider)
	}
}

// Example handler for GitHub operations
func handleGitHub(github GitHub) error {
	fmt.Printf("Processing GitHub repo: %s/%s/%s\n",
		github.Username,
		github.Organization,
		github.Repo)
	// Add your GitHub-specific logic here
	return nil
}

// Example handler for GitLab operations
func handleGitLab(gitlab GitLab) error {
	fmt.Printf("Processing GitLab repo: %s/%s/%s\n",
		gitlab.Username,
		gitlab.Group,
		gitlab.Repo)
	// Add your GitLab-specific logic here
	return nil
}

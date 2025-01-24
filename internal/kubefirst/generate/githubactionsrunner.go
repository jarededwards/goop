package generate

import (
	"fmt"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst/config"
	githubactionsrunner "github.com/jarededwards/goop/internal/kubefirst/github-actions-runner"
	"github.com/jarededwards/goop/internal/utils"
)

func GitHubActionsRunner(cfg *config.Config, path string) error {

	fmt.Printf("building %s Application wrapper\n", githubactionsrunner.Name)

	applicationData := config.ApplicationInfo{
		ChartInfo:              githubactionsrunner.ChartInfo,
		GitopsRepoURL:          cfg.GitopsConfig.RepoURL,
		SyncWave:               40,
		Project:                "default",
		Name:                   githubactionsrunner.Name,
		Namespace:              "argocd",
		DestinationClusterName: "in-cluster",
		ClusterName:            cfg.ClusterName,
	}

	err := buildApplication("argocd/application-wrapper.yaml.tmpl", filepath.Join(path, fmt.Sprintf("%s.yaml", githubactionsrunner.Name)), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	fmt.Printf("building %s Application\n", githubactionsrunner.Name)

	applicationData.SyncWave = 10

	path = filepath.Join(path, "components", githubactionsrunner.Name)

	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	fmt.Printf("building %s helm values\n", githubactionsrunner.Name)

	err = githubactionsrunner.BuildHelmValues(fmt.Sprintf("%s/values.yaml.tmpl", githubactionsrunner.Name), path)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	err = buildApplication("argocd/application.yaml.tmpl", filepath.Join(path, "application.yaml"), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	return nil
}

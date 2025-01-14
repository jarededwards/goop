package generate

import (
	"fmt"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst/config"
	"github.com/jarededwards/goop/internal/kubefirst/reloader"
	"github.com/jarededwards/goop/internal/utils"
)

func Reloader(cfg *config.Config, path string) error {

	fmt.Printf("building %s Application wrapper\n", reloader.Name)

	applicationData := config.ApplicationInfo{
		ChartInfo:              reloader.ChartInfo,
		GitopsRepoURL:          cfg.GitopsConfig.RepoURL,
		SyncWave:               40,
		Project:                "default",
		Name:                   reloader.Name,
		Namespace:              "argocd",
		DestinationClusterName: "in-cluster",
		ClusterName:            cfg.ClusterName,
	}

	err := buildApplication("argocd/application-wrapper.yaml.tmpl", filepath.Join(path, fmt.Sprintf("%s.yaml", reloader.Name)), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	fmt.Printf("building %s Application\n", reloader.Name)

	applicationData.SyncWave = 10

	path = filepath.Join(path, "components", reloader.Name)

	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	fmt.Printf("building %s helm values\n", reloader.Name)

	err = reloader.BuildHelmValues(fmt.Sprintf("%s/values.yaml.tmpl", reloader.Name), path)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	err = buildApplication("argocd/application.yaml.tmpl", filepath.Join(path, "application.yaml"), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	return nil
}

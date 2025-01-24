package generate

import (
	"fmt"
	"path/filepath"

	certmanager "github.com/jarededwards/goop/internal/kubefirst/cert-manager"
	"github.com/jarededwards/goop/internal/kubefirst/config"
	"github.com/jarededwards/goop/internal/utils"
)

func CertManager(cfg *config.Config, path string) error {

	fmt.Printf("building %s Application wrapper\n", certmanager.Name)

	applicationData := config.ApplicationInfo{
		ChartInfo:              certmanager.ChartInfo,
		GitopsRepoURL:          cfg.GitopsConfig.RepoURL,
		SyncWave:               10,
		Project:                "default",
		Name:                   certmanager.Name,
		Namespace:              "argocd",
		DestinationClusterName: "in-cluster",
		ClusterName:            cfg.ClusterName,
	}

	err := buildApplication("argocd/application-wrapper.yaml.tmpl", filepath.Join(path, fmt.Sprintf("%s.yaml", certmanager.Name)), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	fmt.Printf("building %s Application\n", certmanager.Name)

	applicationData.SyncWave = 10

	path = filepath.Join(path, "components", certmanager.Name)

	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	fmt.Printf("building %s helm values\n", certmanager.Name)

	err = certmanager.BuildHelmValues(fmt.Sprintf("%s/values.yaml.tmpl", certmanager.Name), path)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	err = buildApplication("argocd/application.yaml.tmpl", filepath.Join(path, "application.yaml"), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	return nil
}

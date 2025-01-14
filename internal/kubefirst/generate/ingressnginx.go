package generate

import (
	"fmt"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst/config"
	ingressnginx "github.com/jarededwards/goop/internal/kubefirst/ingress-nginx"
	"github.com/jarededwards/goop/internal/utils"
)

func IngressNginx(cfg *config.Config, path string) error {

	fmt.Printf("building %s Application wrapper\n", ingressnginx.Name)

	applicationData := config.ApplicationInfo{
		ChartInfo:              ingressnginx.ChartInfo,
		GitopsRepoURL:          cfg.GitopsConfig.RepoURL,
		SyncWave:               40,
		Project:                "default",
		Name:                   ingressnginx.Name,
		Namespace:              "argocd",
		DestinationClusterName: "in-cluster",
		ClusterName:            cfg.ClusterName,
	}

	err := buildApplication("argocd/application-wrapper.yaml.tmpl", filepath.Join(path, fmt.Sprintf("%s.yaml", ingressnginx.Name)), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	fmt.Printf("building %s Application\n", ingressnginx.Name)

	applicationData.SyncWave = 10

	path = filepath.Join(path, "components", ingressnginx.Name)

	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	fmt.Printf("building %s helm values\n", ingressnginx.Name)

	err = ingressnginx.BuildHelmValues(fmt.Sprintf("%s/values.yaml.tmpl", ingressnginx.Name), path)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	err = buildApplication("argocd/application.yaml.tmpl", filepath.Join(path, "application.yaml"), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	return nil
}

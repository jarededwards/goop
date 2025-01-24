package generate

import (
	"fmt"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst/config"
)

func Registry(cfg *config.Config, path string) error {

	fmt.Printf("building %s Application wrapper\n", "registry")

	name := "registry"

	var chartInfo = config.ChartInfo{
		Name:           name,
		TargetRevision: "HEAD",
	}

	applicationData := config.ApplicationInfo{
		ChartInfo:              chartInfo,
		GitopsRepoURL:          cfg.GitopsConfig.RepoURL,
		SyncWave:               100,
		Project:                "default",
		Name:                   name,
		Namespace:              "argocd",
		DestinationClusterName: "in-cluster",
		ClusterName:            cfg.ClusterName,
	}

	err := buildApplication(fmt.Sprintf("argocd/%s.yaml.tmpl", name), filepath.Join(path, "registry.yaml"), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	return nil

}

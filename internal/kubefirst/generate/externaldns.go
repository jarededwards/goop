package generate

import (
	"fmt"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst/config"
	externaldns "github.com/jarededwards/goop/internal/kubefirst/external-dns"
	"github.com/jarededwards/goop/internal/utils"
)

func ExternalDNS(cfg *config.Config, path string) error {

	fmt.Printf("building %s Application wrapper\n", externaldns.Name)

	applicationData := config.ApplicationInfo{
		ChartInfo:              externaldns.ChartInfo,
		GitopsRepoURL:          cfg.GitopsConfig.RepoURL,
		SyncWave:               40,
		Project:                "default",
		Name:                   externaldns.Name,
		Namespace:              "argocd",
		DestinationClusterName: "in-cluster",
		ClusterName:            cfg.ClusterName,
	}

	err := buildApplication("argocd/application-wrapper.yaml.tmpl", filepath.Join(path, fmt.Sprintf("%s.yaml", externaldns.Name)), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	fmt.Printf("building %s Application\n", externaldns.Name)

	applicationData.SyncWave = 10

	path = filepath.Join(path, "components", externaldns.Name)

	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	fmt.Printf("building %s helm values\n", externaldns.Name)

	helmData := externaldns.ExternalDNSHelmValues{
		ClusterName:        cfg.ClusterName,
		CloudProvider:      cfg.CloudProvider,
		DomainName:         cfg.DomainName,
		Auth:               externaldns.GetAuth(*cfg),
		Provider:           string(cfg.DNS.Provider),
		AuthFromAnnotation: []string{string(config.DNSProviderAWS), string(config.DNSProviderAzure), string(config.DNSProviderGoogle)},
	}

	err = externaldns.BuildHelmValues(fmt.Sprintf("%s/values.yaml.tmpl", externaldns.Name), path, helmData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	err = buildApplication("argocd/application.yaml.tmpl", filepath.Join(path, "application.yaml"), applicationData)
	if err != nil {
		return fmt.Errorf("error building application: %w", err)
	}

	return nil
}

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jarededwards/goop/internal/argocd"
	"github.com/jarededwards/goop/internal/helm"
	"github.com/jarededwards/goop/internal/kubefirst/config"
	externaldns "github.com/jarededwards/goop/internal/kubefirst/external-dns"
	"github.com/jarededwards/goop/internal/utils"

	"sigs.k8s.io/yaml"
)

func main() {

	config, err := config.ReadPlatformConfig()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)

	}

	baseApp, err := argocd.CreateBaseApplication(*config, externaldns.ExternalDNSChartInfo)
	if err != nil {
		fmt.Printf("Error creating application: %v", err)
		os.Exit(1)
	}
	yAppl, err := yaml.Marshal(baseApp)
	if err != nil {
		fmt.Printf("Error marshaling to YAML: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(yAppl))

	path := filepath.Join("gitops/registry/clusters", config.ClusterName, "components", "external-dns")
	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		fmt.Printf("failed to create directory: %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(filepath.Join(path, "application.yaml"), yAppl, 0644)
	if err != nil {
		fmt.Printf("failed to write file: %v", err)
		os.Exit(1)
	}

	values, err := helm.BuildExternalDNSHelmValues(*config)
	if err != nil {
		fmt.Printf("Error building Helm values: %v\n", err)
		os.Exit(1)
	}
	err = os.WriteFile(filepath.Join(path, "values.yaml"), []byte(values), 0644)
	if err != nil {
		fmt.Printf("failed to write file: %v", err)
		os.Exit(1)
	}

	//! decorate annotations for ingress
	//! decorate annotations for sync wave ordering
	//! should add a kubefirst.konstruct.io/$somehting annotation for tracking?

}

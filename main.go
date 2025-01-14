package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst/config"
	"github.com/jarededwards/goop/internal/kubefirst/generate"
	"github.com/jarededwards/goop/internal/utils"
)

func main() {

	cfg, err := config.ReadPlatformConfig()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}

	path := filepath.Join("gitops", "registry", "clusters", cfg.ClusterName)

	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		fmt.Printf("failed to create directory: %v", err)
	}

	err = generate.ExternalDNS(cfg, path)
	if err != nil {
		fmt.Printf("Error generating external-dns app: %v\n", err)
		os.Exit(1)
	}

	err = generate.CertManager(cfg, path)
	if err != nil {
		fmt.Printf("Error generating cert-manager app: %v\n", err)
		os.Exit(1)
	}

	err = generate.IngressNginx(cfg, path)
	if err != nil {
		fmt.Printf("Error generating ingress-nginx app: %v\n", err)
		os.Exit(1)
	}

	//! decorate annotations for ingress
	//! decorate annotations for sync wave ordering
	//! should add a kubefirst.konstruct.io/$somehting annotation for tracking?

}

package certmanager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst"
	"github.com/jarededwards/goop/internal/kubefirst/config"
)

const Name = "cert-manager"

var ChartInfo = config.ChartInfo{
	Name:           Name,
	RepoURL:        "https://charts.jetstack.io",
	TargetRevision: "v1.15.2",
}

func BuildHelmValues(readPath, writePath string) error {

	file, err := kubefirst.CertManager.ReadFile(readPath)
	if err != nil {
		return fmt.Errorf("error reading templates file: %w", err)
	}

	err = os.WriteFile(filepath.Join(writePath, "values.yaml"), []byte(file), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

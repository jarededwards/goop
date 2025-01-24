package githubactionsrunner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jarededwards/goop/internal/kubefirst"
	"github.com/jarededwards/goop/internal/kubefirst/config"
)

const Name = "github-actions-runner"

var ChartInfo = config.ChartInfo{
	Name:           Name,
	RepoURL:        "https://actions-runner-controller.github.io/actions-runner-controller",
	TargetRevision: "0.23.7",
}

func BuildHelmValues(readPath, writePath string) error {
	file, err := kubefirst.GitHubActionsRunner.ReadFile(readPath)
	if err != nil {
		return fmt.Errorf("error reading templates file: %w", err)
	}

	err = os.WriteFile(filepath.Join(writePath, "values.yaml"), []byte(file), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

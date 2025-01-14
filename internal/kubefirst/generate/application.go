package generate

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/jarededwards/goop/internal/kubefirst"
	"github.com/jarededwards/goop/internal/kubefirst/config"
)

func buildApplication(readPath, writePath string, data config.ApplicationInfo) error {

	file, err := kubefirst.ArgoCD.ReadFile(readPath)
	if err != nil {
		return fmt.Errorf("error reading templates file: %w", err)
	}

	tmpl, err := template.New("tmpl").Funcs(config.Funcs).Parse(string(file))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}
	var buff bytes.Buffer

	err = tmpl.Execute(&buff, data)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	fmt.Printf("%+v", buff.String())

	err = os.WriteFile(writePath, buff.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

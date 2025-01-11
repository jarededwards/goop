package helm

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/jarededwards/goop/internal/kubefirst"
	"github.com/jarededwards/goop/internal/kubefirst/config"
	externaldns "github.com/jarededwards/goop/internal/kubefirst/external-dns"
)

// type BuildHelm interface {
// 	BuildHelmValues(values config.ExternalDNSHelmValues) error
// }

func BuildExternalDNSHelmValues(cfg config.Config) (string, error) {
	fmt.Println("building helm values")
	file, err := kubefirst.ExternalDNS.ReadFile(filepath.Join(cfg.ExternalDNS.ExternalDNSHelmChartInfo.Name, "external-dns/values.yaml.tmpl"))
	if err != nil {
		return "", fmt.Errorf("error reading templates file: %w", err)
	}

	tmpl, err := template.New("tmpl").Parse(string(file))
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	data := externaldns.ExternalDNSHelmValues{
		CloudProvider: cfg.CloudProvider,
		DomainName:    cfg.DomainName,
		EnvName:       externaldns.GetExternalDNSAuth(cfg.CloudProvider),
		Provider:      cfg.ExternalDNS.Provider,
	}

	var outputBuffer bytes.Buffer

	err = tmpl.Execute(&outputBuffer, data)
	if err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}
	fmt.Printf("%+v", outputBuffer.String())
	return outputBuffer.String(), nil
}

package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/jarededwards/goop/internal/argocd"
	"github.com/jarededwards/goop/internal/kubefirst"
	"github.com/jarededwards/goop/internal/kubefirst/config"
	externaldns "github.com/jarededwards/goop/internal/kubefirst/external-dns"
	"github.com/jarededwards/goop/internal/utils"
	"sigs.k8s.io/yaml"
)

const externalDNS = "external-dns"

func GenerateExternalDNSApp(config *config.Config) error {

	//! establish the base Application for external dns
	baseApp, err := argocd.CreateBaseApplication(*config, externaldns.ExternalDNSChartInfo)
	if err != nil {
		fmt.Printf("Error creating application: %v", err)
		os.Exit(1)
	}
	// Convert to clean YAML (without status field) in one block
	cleanYAML, err := func(app *v1alpha1.Application) ([]byte, error) {
		yamlData, err := yaml.Marshal(app)
		if err != nil {
			return nil, fmt.Errorf("error marshaling application: %w", err)
		}

		// Convert to map and remove status
		var appMap map[string]interface{}
		if err := yaml.Unmarshal(yamlData, &appMap); err != nil {
			return nil, fmt.Errorf("error unmarshaling to map: %w", err)
		}
		delete(appMap, "status")

		// Marshal back to YAML
		return yaml.Marshal(appMap)
	}(baseApp)
	if err != nil {
		fmt.Printf("error generating clean yaml: %v", err)
		os.Exit(1)
	}

	//! debug
	fmt.Println(string(cleanYAML))

	path := filepath.Join("gitops/registry/clusters", config.ClusterName, "components", externalDNS)
	err = utils.CreateDirIfNotExist(path)
	if err != nil {
		fmt.Printf("failed to create directory: %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(filepath.Join(path, "application.yaml"), cleanYAML, 0644)
	if err != nil {
		fmt.Printf("failed to write file: %v", err)
		os.Exit(1)
	}

	values, err := buildExternalDNSHelmValues(*config)
	if err != nil {
		fmt.Printf("Error building Helm values: %v\n", err)
		os.Exit(1)
	}
	err = os.WriteFile(filepath.Join(path, "values.yaml"), []byte(values), 0644)
	if err != nil {
		fmt.Printf("failed to write file: %v", err)
		os.Exit(1)
	}

	return nil
}

func buildExternalDNSHelmValues(cfg config.Config) (string, error) {
	fmt.Println("building helm values")
	file, err := kubefirst.ExternalDNS.ReadFile(filepath.Join(cfg.DNS.ExternalDNSHelmChartInfo.Name, "external-dns/values.yaml.tmpl"))
	if err != nil {
		return "", fmt.Errorf("error reading templates file: %w", err)
	}

	data := externaldns.ExternalDNSHelmValues{
		CloudProvider:      cfg.CloudProvider,
		DomainName:         cfg.DomainName,
		EnvName:            externaldns.GetAuth(cfg.DNS.Provider),
		Provider:           string(cfg.DNS.Provider),
		AuthFromAnnotation: []string{string(config.DNSProviderAWS), string(config.DNSProviderAzure), string(config.DNSProviderGoogle)},
	}

	var outputBuffer bytes.Buffer

	tmpl, err := template.New("tmpl").Funcs(template.FuncMap{
		"in": func(slice []string, item string) bool {
			for _, s := range slice {
				if s == item {
					return true
				}
			}
			return false
		},
		"notIn": func(slice []string, item string) bool {
			for _, s := range slice {
				if s == item {
					return false
				}
			}
			return true
		},
	}).Parse(string(file))
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}

	err = tmpl.Execute(&outputBuffer, data)
	if err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}

	fmt.Printf("%+v", outputBuffer.String())

	return outputBuffer.String(), nil
}

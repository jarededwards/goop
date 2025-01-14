package externaldns

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jarededwards/goop/internal/kubefirst"
	"github.com/jarededwards/goop/internal/kubefirst/config"
)

const Name = "external-dns"

var ChartInfo = config.ChartInfo{
	Name:           Name,
	RepoURL:        "https://kubernetes-sigs.github.io/external-dns",
	TargetRevision: "1.14.4",
}

type ExternalDNSHelmValues struct {
	ClusterName        string
	CloudProvider      string
	DomainName         string
	Auth               string
	Provider           string
	AuthFromAnnotation []string
}

func BuildHelmValues(readPath, writePath string, data ExternalDNSHelmValues) error {

	file, err := kubefirst.ExternalDNS.ReadFile(readPath)
	if err != nil {
		return fmt.Errorf("error reading templates file: %w", err)
	}

	var buff bytes.Buffer

	tmpl, err := template.New("tmpl").Funcs(config.Funcs).Parse(string(file))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	err = tmpl.Execute(&buff, data)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	fmt.Printf("%+v", buff.String())

	err = os.WriteFile(filepath.Join(writePath, "values.yaml"), buff.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func GetAuth(cfg config.Config) string {

	switch cfg.DNS.Provider {
	case config.DNSProviderCloudflare:
		return "CF_API_TOKEN"
	case config.DNSProviderAkamai:
		return "LINODE_TOKEN"
	case config.DNSProviderAWS:
		return fmt.Sprintf("eks.amazonaws.com/role-arn: arn:aws:iam::%s:role/external-dns-%s", cfg.Cloud.AWS.AccountID, cfg.ClusterName)
	case config.DNSProviderAzure:
		return fmt.Sprintf("azure.workload.identity/client-id: %s", cfg.Cloud.Azure.IdentityClientID)
	case config.DNSProviderCivo:
		return "CIVO_TOKEN"
	case config.DNSProviderDigitalOcean:
		return "DO_TOKEN"
	case config.DNSProviderGoogle:
		return fmt.Sprintf("iam.gke.io/gcp-service-account: external-dns-%s@%s.iam.gserviceaccount.com", cfg.ClusterName, cfg.Cloud.Google.ProjectName)
	case config.DNSProviderVultr:
		return "VULTR_API_KEY"
	default:
		return "NOT_SUPPORTED"
	}
}

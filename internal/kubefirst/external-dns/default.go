package externaldns

import (
	"fmt"

	"github.com/jarededwards/goop/internal/kubefirst/config"
)

var ExternalDNSChartInfo = config.ChartInfo{
	Name:                        "external-dns",
	Namespace:                   "argocd",
	HelmChartRepoURL:            "https://kubernetes-sigs.github.io/external-dns",
	TargetRevision:              "1.14.4",
	HelmChart:                   "external-dns",
	DestinationClusterNamespace: "external-dns",
	DestinationClusterName:      "in-cluster",
	Project:                     "default",
	Annotations: map[string]string{
		"argocd.argoproj.io/sync-wave": "10",
	},
}

type ExternalDNSHelmValues struct {
	ClusterName        string
	CloudProvider      string
	DomainName         string
	Auth               string
	Provider           string
	AuthFromAnnotation []string
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

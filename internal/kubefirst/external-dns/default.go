package externaldns

import "github.com/jarededwards/goop/internal/kubefirst/config"

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
	CloudProvider      string
	DomainName         string
	EnvName            string
	Provider           string
	AuthFromAnnotation []string
}

func GetAuth(dnsProvider config.DNSProvider) string {

	switch dnsProvider {
	case config.DNSProviderCloudflare:
		return "CF_API_TOKEN"
	case config.DNSProviderAkamai:
		return "LINODE_TOKEN"
	case config.DNSProviderAWS:
		return "eks.amazonaws.com/role-arn: arn:aws:iam::<AWS_ACCOUNT_ID>:role/external-dns-<CLUSTER_NAME>"
	case config.DNSProviderAzure:
		return "azure.workload.identity/client-id: <IDENTITY_CLIENT_ID>"
	case config.DNSProviderCivo:
		return "CIVO_TOKEN"
	case config.DNSProviderDigitalOcean:
		return "DO_TOKEN"
	case config.DNSProviderGoogle:
		return "iam.gke.io/gcp-service-account: external-dns-<CLUSTER_NAME>@<GOOGLE_PROJECT>.iam.gserviceaccount.com"
	case config.DNSProviderVultr:
		return "VULTR_API_KEY"
	default:
		return "NOT_SUPPORTED"
	}
}

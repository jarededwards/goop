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
	CloudProvider string
	DomainName    string
	EnvName       string
	Provider      string
}

func GetExternalDNSAuth(cloudProvider string) string {

	switch cloudProvider {
	case "cloudflare":
		return "CF_API_TOKEN"
	case "akamai":
		return "LINODE_TOKEN"
	case "aws":
		return "eks.amazonaws.com/role-arn: arn:aws:iam::<AWS_ACCOUNT_ID>:role/external-dns-<CLUSTER_NAME>"
	case "azure":
		return "azure.workload.identity/client-id: <IDENTITY_CLIENT_ID>"
	case "civo":
		return "CIVO_TOKEN"
	case "digitalocean":
		return "DO_TOKEN"
	case "gcp":
		return "iam.gke.io/gcp-service-account: external-dns-<CLUSTER_NAME>@<GOOGLE_PROJECT>.iam.gserviceaccount.com"
	case "vultr":
		return "VULTR_API_KEY"
	default:
		return "NOT_SUPPORTED"
	}
}

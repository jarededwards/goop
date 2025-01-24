package kubefirst

import (
	"embed"
)

//go:embed argocd/*.yaml.tmpl
var ArgoCD embed.FS

//go:embed cert-manager/*.yaml.tmpl
var CertManager embed.FS

//go:embed external-dns/*.yaml.tmpl
var ExternalDNS embed.FS

//go:embed ingress-nginx/*.yaml.tmpl
var IngressNginx embed.FS

//go:embed reloader/*.yaml.tmpl
var Reloader embed.FS

//go:embed github-actions-runner/*.yaml.tmpl
var GitHubActionsRunner embed.FS

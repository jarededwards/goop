package kubefirst

import (
	"embed"
)

//go:embed external-dns/*.yaml.tmpl
var ExternalDNS embed.FS

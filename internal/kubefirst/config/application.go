package config

// ! deprecated
type BaseApplication struct {
	Annotations                 map[string]string
	DestinationClusterNamespace string
	DestinationClusterName      string
	HelmChart                   string
	HelmChartRepoURL            string
	Name                        string
	Namespace                   string
	Project                     string
	TargetRevision              string
}

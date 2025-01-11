package config

type ChartInfo struct {
	Name                        string
	Namespace                   string
	HelmChartRepoURL            string
	TargetRevision              string
	HelmChart                   string
	DestinationClusterNamespace string
	DestinationClusterName      string
	Project                     string
	Annotations                 map[string]string
}

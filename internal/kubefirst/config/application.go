package config

type ChartInfo struct {
	Name           string
	RepoURL        string
	TargetRevision string
}

type ApplicationInfo struct {
	ChartInfo              ChartInfo
	GitopsRepoURL          string
	SyncWave               int
	Project                string
	Name                   string
	Namespace              string
	DestinationClusterName string
	ClusterName            string
}

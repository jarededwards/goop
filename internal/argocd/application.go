package argocd

import (
	"path/filepath"

	appv1 "github.com/argoproj/argo-cd/v2/pkg/apis/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/jarededwards/goop/internal/kubefirst/config"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateBaseApplication(cfg config.Config, a config.ChartInfo) (*v1alpha1.Application, error) {

	app := v1alpha1.Application{
		TypeMeta: metav1.TypeMeta{
			Kind:       appv1.ApplicationKind,
			APIVersion: v1alpha1.SchemeGroupVersion.Group + "/" + v1alpha1.SchemeGroupVersion.Version,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        a.Name,
			Namespace:   a.Namespace,
			Annotations: a.Annotations,
		},
		Spec: v1alpha1.ApplicationSpec{
			Project: a.Project,
			Sources: []v1alpha1.ApplicationSource{
				{
					RepoURL:        a.HelmChartRepoURL,
					TargetRevision: a.TargetRevision,
					Chart:          a.HelmChart,
					Helm: &v1alpha1.ApplicationSourceHelm{
						ValueFiles:  []string{filepath.Join("$values/registry/clusters", cfg.ClusterName, "components", a.Name, "values.yaml")},
						ReleaseName: a.Name,
					},
				},
				{
					RepoURL:        cfg.GitopsConfig.RepoURL,
					TargetRevision: "HEAD",
					Ref:            "values",
				},
			},
			Destination: v1alpha1.ApplicationDestination{
				Namespace: a.DestinationClusterNamespace,
				Name:      a.DestinationClusterName,
			},
			SyncPolicy: &v1alpha1.SyncPolicy{
				Automated: &v1alpha1.SyncPolicyAutomated{
					Prune:    true,
					SelfHeal: true,
				},
				SyncOptions: []string{
					"CreateNamespace=true",
				},
			},
		},
		Status: v1alpha1.ApplicationStatus{},
	}
	return &app, nil
}

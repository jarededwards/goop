# todo 

get the templates with values yaml in line and turn into go templates
start with k3s
apply command that bootstraps github and applies to the provided kubeconfig 


kubectl kustomize https://github.com/konstructio/manifests/argocd/cloud\?ref\=v1.1.0 | kubectl apply -f -

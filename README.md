# Civo with cloudflare origin certs

### prerequisites
- k3d
- kubectl
- linkerd
- civo token
- cloudflare token
- cloudflare origin ca key
- dns

### clone the `goop` git repository
```sh
git clone .git
cd goop
```

### create a local `k3d` cluster to provision cloud infrastructure with gitops
We'll provision a local k3d cluster that will need a `CIVO_TOKEN` added as a kubernetes secret. This `k3d` cluster will also have a few additional [manifests](../manifests/bootstrap-k3d.yaml) that install argocd to the new cluster with a few default configurations we'll take advantage of.
```sh
k3d cluster create goop --agents "1" --agents-memory "4096m" \
    --volume $PWD/manifests/bootstrap-k3d.yaml:/var/lib/rancher/k3s/server/manifests/bootstrap-k3d.yaml
```

### add your `CIVO_TOKEN`, `CLOUDFLARE_API_TOKEN`,  and `CLOUDFLARE_ORIGIN_CA_KEY` for provisioning cloud infrastructure and managing DNS
The `CIVO_TOKEN` will be used by the crossplane terraform provider to allow for provisioning of CIVO cloud infrastructure as well as for external-dns to create and adjust DNS records in your CIVO cloud account. The `CLOUDFLARE_API_TOKEN` will be used to manage DNS records in your Cloudflare zone and `CLOUDFLARE_ORIGIN_CA_KEY` will be used by the Cloudflare Origin CA Issuer controller to get certificates for TLS communication of the metaphor service.
```sh

# check envs
echo $CIVO_TOKEN
echo $CIVO_TOKEN
echo $AWS_ACCESS_KEY_ID
echo $AWS_SECRET_ACCESS_KEY
echo $CLOUDFLARE_API_TOKEN
echo $CLOUDFLARE_ORIGIN_CA_KEY

kubectl -n crossplane-system create secret generic crossplane-secrets \
  --from-literal=AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
  --from-literal=AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
  --from-literal=CIVO_TOKEN=$CIVO_TOKEN \
  --from-literal=TF_VAR_civo_token=$CIVO_TOKEN \
  --from-literal=TF_VAR_cloudflare_api_token=$CLOUDFLARE_API_TOKEN \
  --from-literal=TF_VAR_cloudflare_origin_issuer_token=$CLOUDFLARE_ORIGIN_CA_KEY
```

### wait for argocd pods in k3d to be running
```sh
watch kubectl get pods -A
```
### get the argocd root password
```sh
kubectl -n argocd get secret/argocd-initial-admin-secret -ojsonpath='{.data.password}' | base64 -D | pbcopy
```
### visit the argocd ui
```sh
kubectl -n argocd port-forward svc/argocd-server 8888:80 
open http://localhost:8888
```

### bootstrap the `k3d` cluster with crossplane and install the terraform provider
```sh
kubectl apply -f https://raw.githubusercontent.com/jarededwards/goop/main/registry/bootstrap/bootstrap.yaml
```

### apply the registry for mgmt to the new remote mgmt cluster argocd
```sh
kubectl apply -f https://raw.githubusercontent.com/jarededwards/goop/main/registry/clusters/mgmt/registry.yaml
```

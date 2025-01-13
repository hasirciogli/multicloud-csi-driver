# Önce Docker image'ı build edip push edin
docker build -t ghcr.io/hasirciogli/multicloud-csi-driver:latest .
docker push ghcr.io/hasirciogli/multicloud-csi-driver:latest

# Sonra Kubernetes manifestini uygulayın
kubectl apply -f kubernetes/deploy.yaml

# Kontrol edin
kubectl get pods -n kube-system | grep multicloud
kubectl get csidriver
kubectl get storageclass
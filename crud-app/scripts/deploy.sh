./scripts/publish-containers.sh
kubectl delete -k .
kubectl apply -k .

docker buildx build --platform linux/amd64,linux/arm64 -t toukop/crud-app:latest .
docker push toukop/crud-app:latest

docker tag toukop/crud-app:latest europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/crud-app:latest

docker push europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/crud-app:latest

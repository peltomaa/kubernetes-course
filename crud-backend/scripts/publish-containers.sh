docker buildx build --platform linux/amd64,linux/arm64 -t toukop/crud-backend:latest .
docker buildx build --platform linux/amd64,linux/arm64 -t toukop/crud-backend-worker:latest -f Dockerfile-worker .
docker push toukop/crud-backend:latest
docker push toukop/crud-backend-worker:latest

docker tag toukop/log-output-worker:latest europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/crud-backend:latest
docker tag toukop/log-output-web:latest europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/crud-backend-worker:latest

docker push europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/crud-backend:latest
docker push europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/crud-backend-worker:latest

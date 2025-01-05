docker buildx build --platform linux/amd64,linux/arm64 -t toukop/log-output-worker:latest -f Dockerfile-worker .
docker buildx build --platform linux/amd64,linux/arm64 -t toukop/log-output-web:latest -f Dockerfile-web .
docker push toukop/log-output-worker:latest
docker push toukop/log-output-web:latest


docker tag toukop/log-output-worker:latest europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/log-output-worker:latest
docker tag toukop/log-output-web:latest europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/log-output-web:latest

docker push europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/log-output-worker:latest
docker push europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/log-output-web:latest

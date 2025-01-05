docker buildx build --platform linux/amd64,linux/arm64 -t toukop/log-output-worker:latest -f Dockerfile-worker .
docker buildx build --platform linux/amd64,linux/arm64 -t toukop/log-output-web:latest -f Dockerfile-web .
docker push toukop/log-output-worker:latest
docker push toukop/log-output-web:latest

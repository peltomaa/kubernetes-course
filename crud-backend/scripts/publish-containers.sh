docker buildx build --platform linux/amd64,linux/arm64 -t toukop/crud-backend:latest .
docker buildx build --platform linux/amd64,linux/arm64 -t toukop/crud-backend-worker:latest -f Dockerfile-worker .
docker push toukop/crud-backend:latest
docker push toukop/crud-backend-worker:latest

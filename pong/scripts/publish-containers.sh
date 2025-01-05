docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t toukop/pong:latest .
docker tag toukop/pong:latest europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/pong:latest
docker push europe-north1-docker.pkg.dev/nomadic-charge-446815-j2/kubernetes-course/pong:latest
docker push toukop/pong:latest

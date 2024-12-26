docker build -t toukop/log-output-worker:latest -f Dockerfile-worker .
docker build -t toukop/log-output-web:latest -f Dockerfile-web .
docker push toukop/log-output-worker:latest
docker push toukop/log-output-web:latest

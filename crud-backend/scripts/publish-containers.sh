docker build -t toukop/crud-backend:latest .
docker build -t toukop/crud-backend-worker:latest -f Dockerfile-worker .
docker push toukop/crud-backend:latest
docker push toukop/crud-backend-worker:latest

# gobby
Cloud native golang project seed, including observability stack

# app
cd app
go build -o gobby
docker build -t gobby . --progress=plain --no-cache
## upload to dockerhub:
docker login
docker tag gobby tomrt/gobby
docker image push tomrt/gobby
## run locally:
docker image ls
docker run -p 9999:9999 gobby

# minikube
## start a cluster
minikube start --nodes 3 -p gobby-cluster
## apply pod only
kubectl apply -f minikube/api-pod.yml

# observability:
    - Prometheus
    - Grafana

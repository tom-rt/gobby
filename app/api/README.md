# API stack
For the API I chose use Huma, which will provide us with automatic OpenAPI spec gen (as well as out of the box API docs). It is router-agnostic which means we can choose our preferred router.
For the router I chose fiber, as it provides some nice abstractions while using fasthttp, the fastest go http library (much faster than the stdlib).
Using zerolog for logging with the fiber zerolog middleware.
For observability, opentelemetry for tracing and metrics. Integrating thanks to the otelfiber middleware.
In the beginning, the DB which will be used is SQLite

# How to deploy and build ?
## API
cd app
go build -o gobby

## Run locally with docker:
colima start
docker build -t gobby . --progress=plain --no-cache
docker run -p 9999:9999 gobby

## Run locally with podman:
podman machine start
podman build -t gobby .
podman run -p 9999:9999 gobby

## Upload to dockerhub:
docker login
docker tag gobby tomrt/gobby
docker image push tomrt/gobby

## Start a minikube cluster
minikube start --nodes 3 -p gobby-cluster

## apply pod only
kubectl apply -f minikube/api-pod.yml

## DONE
Setup FIBER
SETUP zerolog

## TODO API:
Setup HUMA
Setup few tests
Setup opentelemetry with otelfiber
Setup html templates generator ?
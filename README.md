# What is gobby ?
This repository contains the source code of gobby, an open source identity and access management solution, meant to be deployed in kubernetes.

# Roadmap
v1 authentication handling: handle CRUD of users
v2 delegated authentication
v3 extended features: HR functionnalities ?

# The stack
For the API I chose use Huma, which will provide us with automatic OpenAPI spec gen (as well as out of the box API docs). It is router-agnostic which means we can choose our preferred router.

For the router I chose fiber, as it provides some nice abstractions while using fasthttp, the fastest go http library (much faster than the stdlib).

Using zerolog for logging with the fiber zerolog middleware.

For observability, opentelemetry for tracing and metrics. Integrating thanks to the otelfiber middleware.


# How to deploy and build ?
## API
cd app
go build -o gobby
docker build -t gobby . --progress=plain --no-cache

## Upload to dockerhub:
docker login
docker tag gobby tomrt/gobby
docker image push tomrt/gobby

## Run locally:
docker image ls
docker run -p 9999:9999 gobby

## Start a minikube cluster
minikube start --nodes 3 -p gobby-cluster

## apply pod only
kubectl apply -f minikube/api-pod.yml


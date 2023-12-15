# gobby
Cloud native golang project seed, including observability stack

# app
cd app
go build -o gobby
docker build -t gobby .
docker image ls
docker run -p 9999:9999 gobby

# observability:
    - Prometheus
    - Grafana

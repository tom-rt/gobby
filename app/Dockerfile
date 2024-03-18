FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /gobby

EXPOSE 9999

CMD [ "/gobby" ]
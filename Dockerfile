FROM golang:1.22.4-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app/

RUN GOOS=linux go build -o /docker-rest-api

CMD ["/docker-rest-api"]

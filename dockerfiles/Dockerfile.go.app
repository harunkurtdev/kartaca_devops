FROM golang:1.17-alpine AS build

LABEL maintainer="Harun KURT <harunkurtdev> (https://github.com/harunkurtdev)"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

EXPOSE 5555
CMD ["go","run","main.go"]

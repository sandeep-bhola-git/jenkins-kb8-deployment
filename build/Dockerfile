FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 7701
CMD ["/docker-gs-ping"]
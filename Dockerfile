#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build -o main main.go

#final stage
FROM alpine:latest
COPY --from=builder /go/src/app/main /main
EXPOSE 3000
ENTRYPOINT /main
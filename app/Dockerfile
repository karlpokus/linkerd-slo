FROM golang:alpine AS builder
RUN apk --no-cache add ca-certificates
WORKDIR /src
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./ip-app

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /src
COPY --from=builder /src/ip-app ./ip-app
EXPOSE 8989
ENTRYPOINT ["/src/ip-app"]

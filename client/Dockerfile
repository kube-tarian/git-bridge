FROM golang:1.18.3 AS builder
WORKDIR /
COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./client

FROM scratch
COPY --from=builder /client client
USER 1101:1101
ENTRYPOINT ["./client"]

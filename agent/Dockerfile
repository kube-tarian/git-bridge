FROM golang:1.18.3 AS builder
WORKDIR /
COPY ./ ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./git-bridge

FROM scratch
COPY --from=builder /git-bridge git-bridge
USER 1101:1101
ENTRYPOINT ["./git-bridge"]

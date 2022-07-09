FROM golang:1.18-alpine as builder

RUN mkdir /workspace

COPY . /workspace

WORKDIR /workspace

RUN CGO_ENABLED=0 go build -o client .

RUN chmod +x /workspace/client

FROM alpine:latest

RUN mkdir /workspace

COPY --from=builder /workspace/client /workspace

CMD ["/workspace/client"]
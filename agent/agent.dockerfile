FROM golang:1.18-alpine as builder

RUN mkdir /workspace

COPY . /workspace

WORKDIR /workspace

RUN CGO_ENABLED=0 go build -o agent .

RUN chmod +x /workspace/agent

FROM alpine:latest

RUN mkdir /workspace

COPY --from=builder /workspace/agent /workspace

CMD ["/workspace/agent"]
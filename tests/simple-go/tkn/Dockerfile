FROM docker.io/library/golang:1.18 AS builder-tkn
WORKDIR /app
COPY . /app
USER root
RUN go build -v -o tkn ./cmd/tkn

FROM cgr.dev/chainguard/static:latest
COPY --from=builder-tkn /app/tkn /tkn

ENV foo=bar
ENTRYPOINT ["/tkn"]

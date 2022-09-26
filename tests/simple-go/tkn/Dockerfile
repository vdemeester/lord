FROM docker.io/library/golang:1.8 AS builder
WORKDIR /app
COPY . /app
RUN go build -v -o /artifacts/tkn ./cmd/tkn

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /artifacts/tkn /tkn
ENTRYPOINT [ "/tkn" ]

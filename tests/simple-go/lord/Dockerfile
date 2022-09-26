FROM cgr.dev/chainguard/go:1.19 AS builder
WORKDIR /app
COPY . /app
RUN go build -v -o /artifacts/lord .

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /artifacts/lord /lord
ENTRYPOINT [ "/lord" ]

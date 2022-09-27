FROM cgr.dev/chainguard/go:1.19 AS builder-lord
WORKDIR /app
COPY . /app
USER root
RUN go build -v -o lord .

FROM cgr.dev/chainguard/static:latest
COPY --from=builder-lord /app/lord /lord
LABEL foo=bar

ENTRYPOINT ["/lord"]

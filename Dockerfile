FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./sudo-azure-operator /sudo-azure-operator

ENTRYPOINT ["/sudo-azure-operator"]

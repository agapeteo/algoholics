FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add --no-cache openssh git

COPY scripts/generate-site.sh / /

ENTRYPOINT ["/generate-site.sh"]

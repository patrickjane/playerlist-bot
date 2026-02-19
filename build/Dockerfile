FROM scratch

WORKDIR /

ARG TARGETARCH
ARG VERSION

LABEL org.opencontainers.image.version=${VERSION}

COPY playerlistbot.linux-${TARGETARCH} /playerlistbot

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/playerlistbot"]

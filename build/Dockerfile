FROM scratch

WORKDIR /

ARG TARGETARCH
ARG VERSION

LABEL org.opencontainers.image.version=${VERSION}

COPY playerlistbot.linux-${TARGETARCH} /playerlistbot

ENTRYPOINT ["/playerlistbot"]

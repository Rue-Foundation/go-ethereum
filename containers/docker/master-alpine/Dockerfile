FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.7 https://github.com/Rue-Foundation/go-rue && \
  (cd go-rue && make grue) && \
  cp go-rue/build/bin/grue /grue && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-rue && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/grue"]

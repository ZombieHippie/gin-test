# Docker image for the Artsy-Host
#
#     cd $GOPATH/src/github.com/ZombieHippie/test-gin/server
#     make deps build
#     docker build --rm=true -t colelawr/artsy-host .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git && \ 
  rm -rf /var/cache/apk/*

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY ./server /usr/src/app/

EXPOSE 8080

CMD ["./server"]

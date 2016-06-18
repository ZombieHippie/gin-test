# Docker image for the Dryclean-Host
#
#     cd $GOPATH/src/github.com/ZombieHippie/test-gin/server
#     make deps build
#     docker build --rm=true -t colelawr/dryclean-host .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git && \ 
  rm -rf /var/cache/apk/*

WORKDIR /usr/app

ADD dryclean-host /

EXPOSE 8080

CMD ["/dryclean-host"]

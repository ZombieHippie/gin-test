FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    nodejs && \
  rm -rf \
    /var/cache/apk/*

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
ADD package.json /usr/src/app/
ADD typings.json /usr/src/app/
RUN npm install --unsafe-perm

# Bundle app source
COPY . /usr/src/app

ENTRYPOINT ["/usr/src/app/node_modules/.bin/ts-node", "/usr/src/app/drone-plugin.0.5.0.ts"]

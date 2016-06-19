FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    nodejs && \
  rm -rf \
    /var/cache/apk/*

# Create app directory
RUN mkdir -p /dryclean
WORKDIR /dryclean

# Install app dependencies
ADD package.json /dryclean/
ADD typings.json /dryclean/
RUN npm install --unsafe-perm

# Bundle app source
COPY . /dryclean

ENTRYPOINT ["/dryclean/node_modules/.bin/ts-node", "/dryclean/drone-plugin.ts"]

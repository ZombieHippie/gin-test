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
COPY package.json /usr/src/app/
RUN npm install

# Bundle app source
COPY . /usr/src/app

ENTRYPOINT ["npm", "start"]

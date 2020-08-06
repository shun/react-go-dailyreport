FROM node:lts-buster-slim

ARG NG_VER=9.1.3
WORKDIR /home/node

RUN apt update \
    && apt install -y \
        git \
        nginx \
        sudo \
        supervisor \
        whois \
    && rm -rf /var/lib/apt/lists/* \
    && gpasswd -a node sudo \
    && yarn global add \
        @angular/cli@${NG_VER} \
    && ng new frontapp --style=styl --routing=true \
    && chown -R node:node /home/node

USER node
WORKDIR /home/node/frontapp

RUN yarn add \
        @angular/animations@${NG_VER} \
        @angular/common@${NG_VER} \
        @angular/core@${NG_VER} \
        @angular/forms@${NG_VER} \
        @angular/platform-browser@${NG_VER} \
        @angular/cdk \
        axios \
        primeicons \
        primeng \
        primeflex \
        rxjs@^6.5.0 \
        tslib@^1.10.0 \
        zone.js@0.10.2 \
    && npm cache clean -f



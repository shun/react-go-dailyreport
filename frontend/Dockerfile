FROM node:lts-buster-slim

WORKDIR /home/node
ENV WORKUSR=node

RUN apt update \
    && apt install -y \
        git \
        nginx \
        sudo \
        supervisor \
        whois \
    && rm -rf /var/lib/apt/lists/* \
    && yarn global add \
         create-react-app \
         webpack

USER node
RUN create-react-app frontapp

WORKDIR /home/node/frontapp
RUN yarn add \
         axios \
         react-router-dom


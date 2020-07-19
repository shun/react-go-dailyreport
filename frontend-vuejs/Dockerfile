FROM node:lts-buster-slim

ARG WORKUSR
ENV WORKUSR=${WORKUSR}
WORKDIR /home/${WORKUSR}

ADD ./.vuerc /home/${WORKUSR}
RUN apt update \
    && apt install -y \
        git \
        nginx \
        sudo \
        supervisor \
        whois \
    && rm -rf /var/lib/apt/lists/* \
    && yarn global add \
        @vue/cli

USER ${WORKUSR}
ARG VUE_CLI_SKIP_DIRTY_GIT_PROMPT=1

RUN vue create frontapp -p mypreset
WORKDIR /home/${WORKUSR}/frontapp


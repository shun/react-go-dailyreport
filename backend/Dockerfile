FROM golang:1.14-buster

ARG DEBIAN_FRONTEND=noninteractive
ARG WORKUSR=${WORKUSR}
ENV APP_ROOT=/home/${WORKUSR}/app
ENV GO11MODULE=on

RUN apt update \
    && apt upgrade -y \
    && apt install -y \
        mariadb-client \
        sudo \
        vim \
        whois \
    && rm -rf /var/lib/apt/lists/* \
    && go get -u github.com/labstack/echo/... \
    && go get -u github.com/jinzhu/gorm \
    && go get -u github.com/go-sql-driver/mysql \
    && useradd -m -p `echo "${WORKUSR}" | mkpasswd -s -m sha-512` -s /bin/bash ${WORKUSR} && gpasswd -a ${WORKUSR} sudo

USER ${WORKUSR}
WORKDIR ${APP_ROOT}
EXPOSE 8080


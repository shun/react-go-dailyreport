#!/bin/bash
set -eu

function wait_db() {
    until mysqladmin ping -h db -P 3306 --silent;
    do
      echo "Waiting for database connection..."
      sleep 5
    done
  echo "!! database is ready."
}

TMPGID=`id -g ${WORKUSR}`
if [ ! "$TMPGID"="$LOCALGID" ]; then
    echo "change GID"
    groupadd -g 11111 tmpgrp
    usermod -g tmpgrp ${WORKUSR}
    groupdel ${WORKUSR}

    groupadd -g $LOCALGID ${WORKUSR}
    usermod -g $LOCALGID ${WORKUSR}
    usermod -u $LOCALUID ${WORKUSR}
    groupdel tmpgrp
    chown -R $LOCALUID:$LOCALGID ${APP_ROOT}
else
    echo "no change GID"
fi

wait_db

exec "$@"


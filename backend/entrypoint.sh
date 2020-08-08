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

wait_db

if [ -n "${TESTING:-}" ]; then
	exec bash
else
	exec "$@"
fi


#!/usr/bin/env bash

if [ -z "$POSTGRES_PORT" ] || [ -z "$POSTGRES_HOST" ] || [ -z "$POSTGRES_DB" ] || [ -z "$POSTGRES_USER" ] || [ -z "$POSTGRES_PASSWORD" ]; then
    echo "Please, specify DB credentials POSTGRES_DB POSTGRES_USER POSTGRES_PASSWORD"
    exit 1;
fi

LIQUIBASE_URL="jdbc:postgresql://$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB"
LIQUIBASE_USERNAME=$POSTGRES_USER
LIQUIBASE_PASSWORD=$POSTGRES_PASSWORD

exec /entrypoint "$@"
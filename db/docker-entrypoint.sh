#!/bin/sh

# database initialization and start postgres server
initdb -D /var/lib/postgresql/data
pg_ctl -D "/var/lib/postgresql/data" -w start

# run migrations
migrate \
  -path="/migrations" \
  -database="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@127.0.0.1:5432/${POSTGRES_DB}?sslmode=disable" \
  up

# stop internal postgres server
pg_ctl -D "/var/lib/postgresql/data" -m fast -w stop

postgres

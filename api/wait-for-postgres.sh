#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
>&2 echo "host: "
>&2 echo $host
shift
cmd="$@"
>&2 echo "cmd: "
>&2 echo $cmd
>&2 echo "POSTGRES_PASSWORD: "
>&2 echo $POSTGRES_PASSWORD

until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "🌥🌥🌥 Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "☀️☀️☀️ Postgres is up - executing command"
exec $cmd
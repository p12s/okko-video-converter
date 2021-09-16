FROM migrate/migrate:4 as migrate

# FROM postgres:alpine3.14
FROM postgres:13.3-alpine

WORKDIR /

COPY --from=migrate /usr/bin/migrate /usr/local/bin/migrate

COPY ./schema /migrations
COPY ./docker-entrypoint.sh /

# RUN apt-get update -y && apt-get upgrade -y && apt-get install bash -y && apt-get install curl -y && apt-get install nano -y
RUN apk update && apk upgrade && apk add bash &&\
    apk add --no-cache curl && apk add lsof && apk add nano

# RUN mkdir -p /var/lib/postgresql/data  &&\
#     chmod 0700 /var/lib/postgresql/data &&\
#     chown -R postgres:postgres /var/lib/postgresql &&\
RUN chmod +x /docker-entrypoint.sh

USER postgres

RUN chmod 0700 /var/lib/postgresql/data &&\
    chown -R postgres:postgres /var/lib/postgresql &&\
    initdb /var/lib/postgresql/data &&\
    echo "host all  all    0.0.0.0/0  trust" >> /var/lib/postgresql/data/pg_hba.conf &&\
    echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf &&\
    pg_ctl start 
    # /usr/local/bin/migrate -path="/migrations" -database="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@127.0.0.1:5432/${POSTGRES_DB}?sslmode=disable" up
    # psql -U ${POSTGRES_USER} -tc "SELECT 1 FROM pg_database WHERE datname = '${POSTGRES_DB}'" | grep -q 1 || psql -U ${POSTGRES_USER} -c "CREATE DATABASE ${POSTGRES_DB}" &&\
    # psql -U ${POSTGRES_USER} -c "ALTER USER ${POSTGRES_USER} WITH PASSWORD '${POSTGRES_PASSWORD}';" &&\
    # chmod +x /docker-entrypoint.sh

#RUN /usr/local/bin/migrate -path="/migrations" -database="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@127.0.0.1:5432/${POSTGRES_DB}?sslmode=disable" up
#RUN /usr/local/bin/migrate -path="/migrations" -database="postgresql://postgres:qwerty@127.0.0.1:5432/postgres?sslmode=disable" up
#ENTRYPOINT ["migrate", "-path=migrations", "-database=postgresql://postgres:qwerty@127.0.0.1:5432/postgres?sslmode=disable", "up"]
EXPOSE 5432
CMD ["/docker-entrypoint.sh"]

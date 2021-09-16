FROM postgres:13.3-alpine

WORKDIR /

RUN apk update && apk upgrade && apk add bash &&\
    apk add --no-cache curl && apk add lsof && apk add nano

USER postgres

RUN chmod 0700 /var/lib/postgresql/data &&\
    chown -R postgres:postgres /var/lib/postgresql &&\
    initdb /var/lib/postgresql/data &&\
    echo "host all  all    0.0.0.0/0  trust" >> /var/lib/postgresql/data/pg_hba.conf &&\
    echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf &&\
    pg_ctl start 

EXPOSE 5432
CMD ["/docker-entrypoint.sh"]

FROM postgres:9.6

ENV POSTGRES_USER gp
ENV POSTGRES_PASSWORD gp
ENV POSTGRES_DB gp

COPY *.sql /docker-entrypoint-initdb.d/

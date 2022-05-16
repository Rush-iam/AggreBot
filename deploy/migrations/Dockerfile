FROM alpine

RUN apk --no-cache add curl
RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG DB_USER
ARG DB_PASSWORD

ENV GOOSE_DBSTRING="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

WORKDIR /migrations
COPY ./migrations ./

CMD goose postgres up
FROM golang:1.18-alpine as build_stage

WORKDIR /build

COPY ./ ./
RUN go mod download && go mod verify
ARG APP_CMD_PATH
RUN go build -o ./app $APP_CMD_PATH

FROM alpine as production
COPY --from=build_stage /build/app /usr/local/bin/

ENTRYPOINT ["app"]
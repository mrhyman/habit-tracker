FROM golang:1.22-alpine3.19 as build

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache \
    && apk add --no-cache tzdata \
    && apk add --no-cache make \
    && apk add --no-cache curl

WORKDIR /bot
COPY . .

RUN make build

FROM scratch

EXPOSE 8080

COPY --from=build /bot/bin/bot-service /bin/bot-service
COPY --from=build /bot/config /config

ENTRYPOINT ["/bin/bot-service", "-config-path", "./config/prod.yml"]

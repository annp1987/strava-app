FROM golang:1.17-alpine3.13 as build
RUN mkdir -p /app
WORKDIR /go/src/strava-app
COPY . .

ARG release
ENV CGO_ENABLED=0 RELEASE=$release GOSUMDB=off
RUN     go clean --modcache && \
        go mod download && \
        go build -o app main.go



FROM alpine:3.13 as prod
WORKDIR /

COPY --from=build /go/src/strava-app/app /usr/local/bin/app

EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/app" ]
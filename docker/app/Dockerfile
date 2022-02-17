# Compile stage
FROM golang:1.17 AS build-env

ADD . /app
WORKDIR /app

RUN go build -o /server

# Final stage
FROM debian:buster

EXPOSE 8081

WORKDIR /
COPY --from=build-env /server /

CMD ["/server"]
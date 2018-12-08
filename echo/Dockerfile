FROM golang AS build

WORKDIR /go/src/
ADD . /go/src/

RUN CGO_ENABLED=0 GOOS=linux go build -o echo-server .

FROM alpine:latest
COPY --from=build /go/src/echo-server /usr/local/bin/echo-server

ENTRYPOINT ["echo-server"]

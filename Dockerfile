FROM golang:alpine AS build-env
LABEL MAINTAINER "Bijan Tavakoli"

ENV GOPATH /go
WORKDIR /go/src
COPY . /go/src/twiss_backend
RUN cd /go/src/twiss_backend && go build .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/twiss_backend/twiss_backend /app
COPY .env /app

EXPOSE 8080

ENTRYPOINT [ "./twiss_backend" ]
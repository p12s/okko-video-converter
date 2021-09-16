FROM golang:1.16.5-buster AS build

#RUN go version
ENV GOPATH=/

WORKDIR /app/

# Get dependancies - will also be cached if we won't change mod/sum
COPY ./go.mod ./go.sum /app/
RUN go mod download

COPY ./ /app/
# libjpeg-dev - для работы библиотеки github.com/kolesa-team/go-webp, использующей  с-шные библиотеки работы с webp
RUN apt-get update && apt-get install libwebp-dev -y
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo -a -o /video-app ./cmd/main.go

#FROM nginx:1.21-alpine
#FROM nginx:latest
#FROM amd64/alpine:3
FROM elswork/cwebp:1.0.0

WORKDIR /


#RUN apk add --no-cache curl
# RUN apt-get update && apt-get upgrade -y && apt-get install -y curl lsof
RUN rm -rf /var/cache/apk && mkdir /var/cache/apk \
  && apk update && apk upgrade && apk add --no-cache --update bash lsof curl libpng-dev libjpeg-turbo-dev giflib-dev tiff-dev autoconf automake make g++ \
  && rm -rf /var/cache/*
# libwebp-dev libwebp libwebp-tools ca-certificates
#ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib:/usr/bin

RUN wget https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-0.5.2.tar.gz && \
	tar -xvzf libwebp-0.5.2.tar.gz && \
	mv libwebp-0.5.2 libwebp && \
	rm libwebp-0.5.2.tar.gz && \
    cd /libwebp && \
	./configure && \
	make && \
	make install && \
	rm -rf libwebp

COPY --from=build /video-app /video-app
COPY ./configs /configs
# TODO эти переменные передавать в secrets
COPY ./.env ./
COPY ./wait-for-postgres.sh ./
RUN mkdir files

# install psql and make wait-for-postgres.sh executable
RUN apk add --no-cache libc6-compat postgresql-client && chmod +x wait-for-postgres.sh video-app

CMD ["./video-app"]

HEALTHCHECK --interval=5s --timeout=3s --start-period=1s CMD curl --fail http://127.0.0.1:8081/health || exit 1

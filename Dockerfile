FROM golang:alpine

RUN apk add --no-cache git \
    && go get -v github.com/darksasori/tvshow

EXPOSE 8080

CMD ["tvshow"]

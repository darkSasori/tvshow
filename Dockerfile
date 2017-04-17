FROM golang:alpine

RUN apk add --no-cache git \
    && go get -v github.com/darksasori/tvshow

EXPOSE 80

CMD ["tvshow"]

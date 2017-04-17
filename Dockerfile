FROM golang:alpine

RUN go get -v -d https://github.com/darkSasori/tvshow.git

CMD ["tvshow"]

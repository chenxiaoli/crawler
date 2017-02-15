FROM golang:1.8-rc

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app

COPY vendors /go/src/
RUN go-wrapper download
RUN go-wrapper install

CMD ["go", "run","./main.go" ,"./config/config.ini"]
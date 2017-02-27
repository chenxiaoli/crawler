FROM golang:1.6

RUN mkdir -p /go/src/github.com/chenxiaoli/crawler
WORKDIR /go/src/github.com/chenxiaoli/crawler
COPY . /go/src/github.com/chenxiaoli/crawler

COPY vendors/ /go/src/
#RUN rm -rf /go/src/app/vendors
#RUN go-wrapper download
#RUN go-wrapper install

CMD ["go", "run","./main.go" ,"./config/config.ini"]
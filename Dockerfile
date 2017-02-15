FROM golang:1.6-onbuild

CMD ["go", "run","*.go  ./config/config.ini"]
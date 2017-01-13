install

go get github.com/chenxiaoli/crawler


[program:crawler]
command=/var/gocode/src/github.com/chenxiaoli/crawler/crawler start-worker /var/gocode/src/github.com/chenxiaoli/crawler/config.ini
autorestart=true
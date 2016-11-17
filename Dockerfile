FROM golang:alpine

ADD . /go/src/github.com/bodva/anonymous
RUN \
 cd /go/src/github.com/bodva/anonymous && \
 go get -v && \
 go build -o /srv/anonymous && \
 rm -rf /go/src/*


WORKDIR /srv

CMD ["/srv/anonymous"]
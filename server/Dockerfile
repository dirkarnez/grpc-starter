FROM golang

COPY again /go/bin

COPY ./src /go/src/github.com/dirkarnez/mentor
WORKDIR /go/src/github.com/dirkarnez/mentor
VOLUME ["/go/src/github.com/dirkarnez/mentor"]

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure
RUN go build

CMD again --bin=main run
EXPOSE 8900
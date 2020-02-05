FROM golang:1.13

ENV GOPATH /go
WORKDIR $GOPATH/src/github.com/faith/alien-invasion
COPY . .

RUN make

CMD ./main -aliens=4

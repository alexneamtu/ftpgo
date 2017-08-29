FROM golang
ADD . $GOPATH/src/ftpgo
RUN go get ./...
RUN go install ./...
EXPOSE 11990
EXPOSE 20000-20010
ENTRYPOINT "./bin/ftpgo"


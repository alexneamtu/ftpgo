FROM golang
ADD . $GOPATH/src/ftpgo
RUN go get ./...
RUN go install ./...
EXPOSE 990
ENTRYPOINT "./bin/ftpgo"


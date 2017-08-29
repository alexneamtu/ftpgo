FROM golang
ADD . $GOPATH/src/ftpgo
RUN go get ./...
RUN go install ./...
RUN mkdir -p /home/test/folders
RUN chmod -R 0777 /home/test
EXPOSE 11990
EXPOSE 20000-20010
ENTRYPOINT "./bin/ftpgo"


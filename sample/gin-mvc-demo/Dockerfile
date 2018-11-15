FROM golang:latest

WORKDIR $GOPATH/src/github.com/Theodoree/sample_project/sample/gin-mvc-demo
COPY .  $GOPATH/src/github.com/Theodoree/sample_project/sample/gin-mvc-demo
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./gin-mvc-demo"]
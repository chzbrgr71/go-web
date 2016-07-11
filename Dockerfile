FROM golang
 
ADD ./src /go/src/go-hello-web
RUN go install go-hello-web
ENTRYPOINT /go/bin/go-hello-web
 
EXPOSE 8001
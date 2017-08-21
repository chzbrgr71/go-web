# build stage
FROM golang:latest AS build-env
MAINTAINER Brian Redmond <brianisrunning@gmail.com>
WORKDIR /app
ADD ./*.go /app/
RUN cd /app && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-web

# final stage
FROM golang:alpine
MAINTAINER Brian Redmond <brianisrunning@gmail.com>
WORKDIR /app
COPY --from=build-env /app/go-web /app/
ENTRYPOINT /app/go-web
EXPOSE 8080
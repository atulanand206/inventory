FROM golang:1.17-alpine

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV PATH=${PATH}:/go/src
ENV APP_HOME go/src/github.com/atulanand206/inventory
RUN mkdir -p ${APP_HOME}
ADD . ${APP_HOME}
WORKDIR ${APP_HOME}
RUN go get -d -v ./...
RUN go mod download
RUN go mod vendor
RUN go mod verify
RUN go build
EXPOSE 5000
CMD [ "go", "run", "main.go" ]
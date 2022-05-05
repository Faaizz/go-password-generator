FROM golang:alpine

COPY . /app
WORKDIR /app

RUN go build -v .
ENTRYPOINT [ "sre-task" ]

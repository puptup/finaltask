FROM golang:1.14.3-alpine3.11

WORKDIR app/

COPY project .

RUN go build server.go

CMD ["./server"]


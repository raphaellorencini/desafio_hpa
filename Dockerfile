FROM golang:1.15rc1-alpine3.12

WORKDIR /go/src/app
COPY . .

RUN go build main.go

EXPOSE 8000

ENTRYPOINT ["./main"]
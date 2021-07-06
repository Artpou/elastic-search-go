FROM golang:1.16-buster

COPY main.go .

RUN go build main.go

EXPOSE 8080

CMD ["./main"]

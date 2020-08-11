FROM golang:latest

RUN mkdir /app

ADD . /app/ 
WORKDIR /app
RUN go build app.go

CMD ["/app/app"]

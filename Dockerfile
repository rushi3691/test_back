FROM golang:1.19.4-alpine

COPY . .

RUN go mod download


RUN go build -o app

EXPOSE 8080

CMD [ "./app" ]
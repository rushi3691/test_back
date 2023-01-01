FROM golang:1.19.4-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
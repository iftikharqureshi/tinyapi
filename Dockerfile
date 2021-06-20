FROM golang:latest

LABEL maintainer="IQ <iftikharqureshi@yahoo.com>"

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build

CMD [ "./tinyapi" ]
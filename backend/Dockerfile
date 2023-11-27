FROM golang:1.21.4-alpine3.18

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app

CMD ["air", "-c", ".air.toml"]
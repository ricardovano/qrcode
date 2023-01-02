FROM golang:1.18.9-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /qrcode

EXPOSE 3000

CMD [ "/qrcode" ]
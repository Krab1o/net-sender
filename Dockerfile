FROM golang:alpine

WORKDIR /app/net-sender

RUN apk add --no-cache git tzdata gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go env -w CGO_ENABLED=1
RUN go build -o ./out/net-sender .

EXPOSE 8080

CMD ["./out/net-sender"]
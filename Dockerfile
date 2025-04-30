FROM golang:1.24-alpine

WORKDIR /app

# Install git and timezone data
RUN apk add --no-cache git tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GO111MODULE=on

RUN go build -o main ./src

CMD ["./main"]

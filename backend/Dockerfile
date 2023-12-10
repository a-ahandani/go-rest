# backend/Dockerfile
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

ENV GOFLAGS="-buildvcs=false"

RUN go install github.com/cosmtrek/air@latest

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["air"]

FROM golang:1.15

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o hgbackend main.go

RUN chmod +x hgbackend

EXPOSE ${APP_PORT}

CMD ["./hgbackend"]





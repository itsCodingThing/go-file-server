FROM golang:alpine

WORKDIR /app/go-file-server

ENV PORT=3001

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go get

EXPOSE 3001

CMD ["go", "run", "server.go"]
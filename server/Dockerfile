FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o server .

EXPOSE ${PORT}

CMD ["./server"]

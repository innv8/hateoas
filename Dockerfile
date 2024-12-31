FROM golang:alpine

WORKDIR /app 
COPY . .
RUN go build -o api -ldflags="-s -w" main.go 

CMD ["./api"]

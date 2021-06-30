ARG VERSION
FROM golang:1.15
WORKDIR /app
ADD . /app/
RUN go build -o ./app/test-ticket-go ./main.go
CMD ["./app/test-ticket-go"]

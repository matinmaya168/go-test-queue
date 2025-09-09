FROM golang:1.24

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o payment-queue
EXPOSE 8081 3316
CMD ["./payment-queue"]
# Payment Queue API

A RESTful API for managing a payment processing queue, built with Go, Gin, MySQL, and Swagger. The API supports enqueuing payments, listing payments, and retrieving payment details, with JWT authentication and rate limiting. Payments are processed in FIFO order using a MySQL-backed queue, with a background worker handling processing.

## Features
- **Endpoints**:
  - `POST /payments`: Enqueue a new payment.
  - `GET /payments`: List all payments.
  - `GET /payments/{id}`: Get a specific payment by ID.
- **Authentication**: JWT-based authentication (Bearer token required).
- **Rate Limiting**: Limits to 10 requests per minute per IP.
- **Persistence**: Stores payments in MySQL with status tracking (`pending`, `processed`, `failed`).
- **Swagger Documentation**: Interactive API docs at `/swagger/index.html`.
- **Logging**: Structured logging with `zerolog` for API and queue operations.
- **Dockerized**: Deployable with Docker and Docker Compose.

## Prerequisites
- **Go**: 1.22+
- **Docker** and **Docker Compose**: For containerized deployment.
- **MySQL**: A running MySQL instance (provided via Docker Compose).
- **swag CLI**: For generating Swagger documentation:
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest

Project Structure
payment-queue/
├── go.mod
├── go.sum
├── main.go
├── Dockerfile
├── docker-compose.yml
├── db/
│   └── db.go
├── models/
│   └── payment.go
├── handlers/
│   └── payment.go
├── middleware/
│   └── auth.go
│   └── ratelimit.go
├── logs/
│   └── app.log
└── docs/
    ├── docs.go
    ├── swagger.json
    ├── swagger.yaml

Setup

Clone the Repository:
git clone <repository-url>
cd payment-queue


Install Dependencies:
go mod tidy


Generate Swagger Docs:
swag init

This generates the docs/ folder with Swagger files.

Configure Environment:

The MySQL connection is defined in db/db.go:dsn := "root:password@tcp(mysql:3306)/paymentdb?charset=utf8mb4&parseTime=True&loc=Local"

Update if using a different MySQL setup.
The JWT secret is in middleware/auth.go:return []byte("your-secret-key")

Replace with a secure key in production (e.g., use environment variables).



Running the Application
With Docker

Start the application and MySQL:
docker-compose up --build

This starts the API at http://localhost:8080 and MySQL at localhost:3306.

Access the Swagger UI:

Open http://localhost:8080/swagger/index.html in your browser.



Without Docker

Start a MySQL server locally (e.g., via Docker):
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=paymentdb mysql


Run the Go application:
swag init
go run main.go



API Usage
Authentication

All endpoints require a JWT token in the Authorization header (Bearer <token>).
Generate a JWT token using HS256 and the secret your-secret-key (e.g., via jwt.io).Example payload:{
  "sub": "user1",
  "exp": 1754659200
}



Endpoints

POST /payments: Enqueue a payment.curl -X POST http://localhost:8080/payments \
-H "Authorization: Bearer <your-jwt-token>" \
-H "Content-Type: application/json" \
-d '{"user_id":"user1","amount":99.99,"product_id":"item1"}'


GET /payments: List all payments.curl -H "Authorization: Bearer <your-jwt-token>" http://localhost:8080/payments


GET /payments/{id}: Get a payment by ID.curl -H "Authorization: Bearer <your-jwt-token>" http://localhost:8080/payments/1



Swagger UI

Access at http://localhost:8080/swagger/index.html.
Click Authorize, enter Bearer <your-jwt-token>, and test endpoints interactively.

Testing

Use the Swagger UI to test endpoints.
Monitor logs in logs/app.log or the console for queue processing:INFO 2025-09-09T17:04:00Z payment enqueued id=1 user_id=user1 amount=99.99 product_id=item1
INFO 2025-09-09T17:04:01Z processing payment id=1 user_id=user1 amount=99.99 product_id=item1
INFO 2025-09-09T17:04:03Z payment processed id=1



Production Notes

Security:
Store the JWT secret in an environment variable or secret manager.
Restrict /swagger/* to authorized users.


Scaling:
For high traffic, consider a dedicated message queue (e.g., RabbitMQ) instead of MySQL polling.
Use Redis for distributed rate limiting.


Monitoring:
Integrate Prometheus/Grafana for metrics on queue size and processing times.


Payment Gateway:
Replace the simulated processing in handlers/payment.go with a real gateway (e.g., Stripe).



Troubleshooting

Swagger UI Not Loading: Ensure swag init was run and docs/ exists.
401 Unauthorized: Verify the JWT token’s validity and secret.
429 Too Many Requests: Wait a minute or increase the rate limit in main.go.
Database Errors: Check MySQL connection details in db/db.go.

Contributing

Fork the repository, make changes, and submit a pull request.
Ensure tests pass and Swagger docs are updated (swag init).

License
MIT License 
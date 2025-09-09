curl -X POST http://localhost:8081/payments \
-H "Authorization: Bearer <your-jwt-token>" \
-H "Content-Type: application/json" \
-d '{"user_id":"user1","amount":99.99,"product_id":"item2", "status": "pending"}'
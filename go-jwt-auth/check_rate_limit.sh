for i in $(seq 1 15); do
  curl -s -o /dev/null -w "req $i: %{http_code}\n" -X POST http://localhost:8080/api/v1/auth/login \
    -H "Content-Type: application/json" \
    -d '{"email":"test@test.com","password":"wrongpassword"}'
done
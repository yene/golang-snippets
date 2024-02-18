# hurl is used to execute HTTP requests. brew install hurl
# hurl test.hurl

GET http://localhost:8080/users


POST http://localhost:8080/users/4
Content-Type: application/json
{
  "email": "john@example.org"
}

# Exia
This is a simple REST API server written in Go, example for learning purposes.

```
git pull git@github.com:keix/exia.git
cd exia
docker compose up --build
curl -X POST http://localhost:8080/users -d '{"Name": "Hello", "Age": 20}' -H "Content-Type: application/json"
```

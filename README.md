# go_ecommerce
https://www.youtube.com/playlist?list=PL5dTjWUk_cPaf5uSEmr8ilR-GtO6s7QJJ

## run

```bash
go run main.go
```

## test

```bash
curl -s localhost:8000/users/productview  | jq .

curl -X POST http://localhost:8000/users/signup \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "太郎",
    "last_name":  "山田",
    "password":   "secret123",
    "email":      "taro@example.com",
    "phone":      "000-0000-0000"
  }' | jq .
```

# go_ecommerce
https://www.youtube.com/playlist?list=PL5dTjWUk_cPaf5uSEmr8ilR-GtO6s7QJJ

## run

```bash
go run main.go
```

## test

```bash
curl -s http://localhost:8000/users/productview  | jq .

curl -s http://localhost:8000/users/search?name=Product  | jq .

curl -sX POST http://localhost:8000/users/addproduct \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Product 1",
    "price": 100,
    "rating": 3
  }' | jq .
```

```bash
curl -sX POST http://localhost:8000/users/signup \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "太郎",
    "last_name":  "山田",
    "password":   "secret123",
    "email":      "taro@example.com",
    "phone":      "000-0000-0000"
  }' | jq .

curl -sX POST http://localhost:8000/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email":      "taro@example.com",
    "password":   "secret123"
  }' | jq .
```

```bash
curl -sX POST http://localhost:8000/addtocart?id=<user_id> \
  -H "Content-Type: application/json" \
  -d '{
    "house_name": "House 1",
    "street_name": "Street 1",
    "city_name": "City 1",
    "pin_code": "000-0000"
  }' | jq .
```

```bash
curl -s http://localhost:8000/listcart?id=<user_id>  | jq .

curl -s "http://localhost:8000/addtocart?id=<product_id>&userID=<user_id>"

curl -s "http://localhost:8000/cartcheckout?id==<user_id>"

curl -s "http://localhost:8000/addtocart?pid=<product_id>&user_id=<user_id>"
```

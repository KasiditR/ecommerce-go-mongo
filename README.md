# Ecommerce With Golang - MongoDB

```bash
docker-compose up -d
go run main.go
```

### POST SIGNUP

http://localhost:8000/users/signup

Request

```json
{
  "first_name": "kasidit",
  "last_name": "ruangmunjit",
  "email": "kasidit.ruangmunjit@gmail.com",
  "password": "123456789",
  "phone": "0924651031"
}
```

Response

```text
Successfully signed in
```

### POST LOGIN

http://localhost:8000/users/login

Request

```json
{
  "email": "kasidit.ruangmunjit@gmail.com",
  "password": "123456789"
}
```

Response

```json
{
  "_id": "67a06ac1f4fdd8e1140dafb3",
  "first_name": "kasidit",
  "last_name": "ruangmunjit",
  "password": "$2a$14$RAzWDCOLgl/F.y0YNO70z.QXwucpC07LNJauEIXb8aSVN2LW5xOsq",
  "email": "kasidit.ruangmunjit@gmail.com",
  "phone": "0924651031",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Imthc2lkaXQucnVhbmdtdW5qaXRAZ21haWwuY29tIiwiRmlyc3RfTmFtZSI6Imthc2lkaXQiLCJMYXN0X05hbWUiOiJydWFuZ211bmppdCIsIlVpZCI6IjY3YTA2YWMxZjRmZGQ4ZTExNDBkYWZiMyIsImV4cCI6MTczODY1NDAyNn0.KOmJSkqrBNPxgB9kFAFZZPKdDj3kC3xOMbJYwcgROuA",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE3MzkxNzI0MjZ9.u2BoY7UAFAmETab4a79Z2O0C9XikERDf5FLnPXVgQGw",
  "created_at": "2025-02-03T07:05:37Z",
  "updated_at": "2025-02-03T07:05:37Z",
  "user_id": "67a06ac1f4fdd8e1140dafb3",
  "usercart": [],
  "address": [],
  "orders": []
}
```

### POST ADD PRODUCT

http://localhost:8000/admin/addproduct

Request

```json
{
  "product_name": "Ipad x15",
  "price": 100,
  "rating": 10,
  "image": "ipad.jpg"
}
```

Response

```text
successfully added
```

### POST ADD ADDRESS

http://localhost:8000/users/addaddress?userID=xxxxxxxxxxxxx

Request

```json
{
  "house_name": "12/34 Bangkok",
  "street_name": "Ramkhamhaeng",
  "city_name": "Bangkok",
  "pin_code": "12345"
}
```

Response

```text
Add address success
```

### PUT EDIT HOME ADDRESS

http://localhost:8000/users/edithomeaddress?userID=xxxxxxxxxxxxx

Request

```json
{
  "house_name": "56/78 Bangkok",
  "street_name": "Seri Thai",
  "city_name": "Bangkok",
  "pin_code": "65482"
}
```

Response

```text
Successfully update the home address
```

### PUT EDIT WORK ADDRESS

http://localhost:8000/users/editworkaddress?userID=xxxxxxxxxxxxx

Request

```json
{
  "house_name": "56/78 Bangkok",
  "street_name": "Ramkhamhaeng",
  "city_name": "Bangkok",
  "pin_code": "54321"
}
```

Response

```text
Successfully update the work address
```

### DELETE DELETE ADDRESS

http://localhost:8000/users/deleteaddress?userID=xxxxxxxxxxxxx

Response

```text
Successfully Deleted
```

### GET USER

http://localhost:8000/users?userID=xxxxxxxxxxxxx

Response

```json
{
  "_id": "67a06ac1f4fdd8e1140dafb3",
  "first_name": "kasidit",
  "last_name": "ruangmunjit",
  "password": "$2a$14$RAzWDCOLgl/F.y0YNO70z.QXwucpC07LNJauEIXb8aSVN2LW5xOsq",
  "email": "kasidit.ruangmunjit@gmail.com",
  "phone": "0924651031",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Imthc2lkaXQucnVhbmdtdW5qaXRAZ21haWwuY29tIiwiRmlyc3RfTmFtZSI6Imthc2lkaXQiLCJMYXN0X05hbWUiOiJydWFuZ211bmppdCIsIlVpZCI6IjY3YTA2YWMxZjRmZGQ4ZTExNDBkYWZiMyIsImV4cCI6MTczODY1OTI2MH0.9SOjemB47bqcb6bkGl9133i_zyCYd7hgpGZ2lOP-Y2I",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE3MzkxNzc2NjB9.yRfq-1OQK-T2F3ls7-BSgf5V3uosk_AKqTZv6VNW5ms",
  "created_at": "2025-02-03T07:05:37Z",
  "updated_at": "2025-02-03T07:05:37Z",
  "user_id": "67a06ac1f4fdd8e1140dafb3",
  "usercart": [],
  "address": [],
  "orders": []
}
```

### GET PRODUCT VIEW

http://localhost:8000/users/productview

Response

```json
[
  {
    "Product_ID": "67a06b64fb81831890f42971",
    "product_Name": "Alienware x15",
    "price": 2500,
    "rating": 10,
    "image": "alienware.jpg"
  },
  {
    "Product_ID": "67a08923e02e02f7d544d1f6",
    "product_Name": "Ipad x15",
    "price": 100,
    "rating": 10,
    "image": "ipad.jpg"
  }
]
```

### GET SEARCH

http://localhost:8000/users/search?name=Ip

Response

```json
[
  {
    "Product_ID": "67a08923e02e02f7d544d1f6",
    "product_Name": "Ipad x15",
    "price": 100,
    "rating": 10,
    "image": "ipad.jpg"
  }
]
```

### GET ADD TO CART

http://localhost:8000/users/addtocart?id=xxxxxxxxxxxxx&userID=xxxxxxxxxxxxx

Response

```text
Successfully
```

### GET REMOVE ITEM

http://localhost:8000/users/removeitem?id=xxxxxxxxxxxxx&userID=xxxxxxxxxxxxx

Response

```text
Successfully
```

### GET CART CHECK OUT

http://localhost:8000/users/cartcheckout?userID=xxxxxxxxxxxxx

Response

```text
Successfully place order
```

### GET INSTANT BUY

http://localhost:8000/users/instantbuy?id=xxxxxxxxxxxxx&userID=xxxxxxxxxxxxx

Response

```text
Successfully
```

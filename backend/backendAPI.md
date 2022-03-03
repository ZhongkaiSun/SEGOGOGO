# Backend API
## Customer API

### 1. Customer Register API

```java
POST    http://localhost:1016/customer/register
```
Create a new customer with the provided information

#### *Example Request*
```java
POST    http://localhost:1016/customer/register
```
Body:
```java
{
    "username":"PIPIKAI",
    "password":"123456",
    "addressLine1":"3824 SW Archer Road",
    "addressLine2":"Apt 5-501",
    "phone":"7839291765",
    "email":"kdlaif@gmail.com"
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6IlN1blpob25nazIyMmFpIiwiZXhwIjoxNjQ2ODU4Njk4LCJpYXQiOjE2NDYyNTM4OTgsImlzcyI6IkdhaW5EYXNoLnRlY2giLCJzdWIiOiJjdXN0b21lciB0b2tlbiJ9.MG6d-gA_FzW2afuqxDXUdGuJY69AiYnE5cizLoRNVp0"
    },
    "msg": "Successfully"
}

```


#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error
- **500**: System error


### 2. Customer Login API

```java
GET    http://localhost:1016/customer/login
```

Login if username and password are correct

#### *Example Request*

```java
GET    http://localhost:1016/customer/login
```

Parameters:

- username:PIPIKAI
- password:123456

#### *Example Response*

```java
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6Ilpob25na2FpU3VuIiwiZXhwIjoxNjQ2ODU5MzI2LCJpYXQiOjE2NDYyNTQ1MjYsImlzcyI6IkdhaW5EYXNoLnRlY2giLCJzdWIiOiJjdXN0b21lciB0b2tlbiJ9.fxlfySWO7evdcKNxzvm_LulnMs_qGYdgOzGYtefM9Ro"
    },
    "msg": "Successfully log in"
}
```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

### 3. Customer Delete API

```java
GET    http://localhost:1016/customer/delete
```

Delete the account if username and password are correct

#### *Example Request*

```java
GET    http://localhost:1016/customer/delete
```

Parameters:

- username:PIPIKAI
- password:123456

#### *Example Response*

```java
{
    "code": 200,
    "data": null,
    "msg": "Successfully deleted"
}
```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

## Rating API

### 1. Rating Create API

```java
POST    http://localhost:1016/rating/create
```
Create a new rating with the provided information

#### *Example Request*
```java
POST    http://localhost:1016/rating/create
```
Body:
```java
{
    "username": "ZhongkaiSun",
    "restaurantName": "Popeyes",
    "star": 5,
    "comment": "Taste so good !!!!",
    "ratingDate": "2022/01/06"
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

### 2. Rating GET API

//


## Cuisine API

### 1. Cuisine Create API

```java
POST    http://localhost:1016/cuisine/create
```
Create a new cuisine for a restaurant

#### *Example Request*
```java
POST    http://localhost:1016/cuisine/create
```
Request Body:
```java
{
"name":"banana pie",        
"restaurantName":"Popeyes", 
"price":2.0,        
"calories":3
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

### 2. Cuisine Read API


```java
GET    localhost:1016/cuisine/read?restaurantName=Popeyes
```
Read a cuisine from a restaurant

#### *Example Request*
```java
GET    http://localhost:1016/cuisine/create
```

#### *Example Response*

```java

{
    "code": 200,
    "data": [
        {
            "name": "8PC Nuggets A La Carte",
            "restaurantName": "Popeyes",
            "price": 4.79,
            "calories": 0
        },
        {
            "name": "apple pie",
            "restaurantName": "Popeyes",
            "price": 2,
            "calories": 3
        },
        {
            "name": "banana pie",
            "restaurantName": "Popeyes",
            "price": 2,
            "calories": 3
        },
        {
            "name": "Mixed Chicken",
            "restaurantName": "Popeyes",
            "price": 22.79,
            "calories": 0
        },
        {
            "name": "Red Beans and Rice",
            "restaurantName": "Popeyes",
            "price": 4.19,
            "calories": 0
        }
    ],
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

### 3. Cuisine Delete API

```java
POST    http://localhost:1016/cuisine/delete
```
Delete the cuisine from a restaurant

#### *Example Request*
```java
POST    http://localhost:1016/cuisine/delete
```
Request Body:
```java
{
"name":"banana pie",        
"restaurantName":"Popeyes", 
"price":2.0,        
"calories":3
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error

## Order API

### 1. Order Create API

```java
POST    http://localhost:1016/order/create
```
Create a new order for a customer

#### *Example Request*
```java
POST    http://localhost:1016/order/create
```
Request Body:
```java
{
    "userName":"Raindrop",
	"restaurantName":"Checkers",
	"orderDate":"03/02/2022",
	"price":8.88,
	"cuisineName":"burger"
}
```
#### *Example Response*

```java

{
    "code": 200,
    "data": null,
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error
  
### 2. Order Read API

```java
GET    http://localhost:1016/order/read?username=Raindrop
```
Create a new cuisine for a restaurant

#### *Example Request*
```java
GET    http://localhost:1016/order/read?username=Raindrop
```

#### *Example Response*

```java

{
    "code": 200,
    "data": [
        {
            "username": "Raindrop",
            "restaurantName": "Checkers",
            "orderDate": "03/02/2022",
            "price": 8.88,
            "cuisineName": "burger"
        }
    ],
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error
- 
## Restaurant API

### 1. Restaurant Read API
```java
GET    http://localhost:1016/restaurant/read
```
Read a restaurant or get the restaurant list 

#### *Example Request*
```java
GET    http://localhost:1016/restaurant/read?name=Popeyes
```
#### *Example Response*

```java

{
    "code": 200,
    "data": {
        "name": "Popeyes",
        "address": "1412N Main St, Gainesville, FL 32601, USA",
        "deliveryFee": 3.99,
        "imgPath": "",
        "typeofMeal": "Chicken",
        "rating": 4.1
    },
    "msg": "Successfully"
}

```

#### *Example Request*
```java
GET    http://localhost:1016/restaurant/read
```
#### *Example Response*

```java

{
    "code": 200,
    "data": [
        {
            "name": "Popeyes",
            "address": "1412N Main St, Gainesville, FL 32601, USA",
            "deliveryFee": 3.99,
            "imgPath": "",
            "typeofMeal": "Chicken",
            "rating": 4.1
        },
        {
            "name": "Checkers",
            "address": "3325 W University Ave, Gainesville, FL 32607, USA",
            "deliveryFee": 1.99,
            "imgPath": "",
            "typeofMeal": "Burgers",
            "rating": 4.2
        }
    ],
    "msg": "Successfully"
}

```

#### *Status Codes*

- **200**: No error
- **422**: Internal Server Error



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



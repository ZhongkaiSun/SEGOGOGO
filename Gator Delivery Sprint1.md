# Gator Delivery Sprint 1



## Author Info

#### Back-end

Zhongkai Sun (sunzhongkai@ufl.edu 7888-5317)

Yudi Zheng (zhengyudi@ufl.edu 6715-2504)

#### Front-end

Yinghuan Zhang (zhangyinghuan@ufl.edu, 2741-2242)

Hongru Liu (hongru.liu@ufl.edu, 5369-7439)



## Project Repository Link

https://github.com/ZhongkaiSun/SEGOGOGO



## Demo Video Link

Demo link of front-end: https://youtu.be/4-t4HFvWU4I

Demo link of back-end:https://www.youtube.com/watch?v=aYVdmD9AhwM

## Accomplishment in Front-end

For our Gator Delivery project, we implemented several pages using **HTML, CSS, JavaScript**, and **Bootstrap 5** in the front-end part:

- Home page: Base of our application, lead to other subpage. Includes functions like searching and paging.
- History page: List previous orders and their details.
- Login page: Users can login to account with username and password.
- Sign-up page: Users can sign-up for an account with name, email, and create their password.
- Account informaiton page: Users can review and change their personal information.
- Payment information page: Users can store and change their card information.
- Change password page: Users can change password if they forget it.



##  Accomplishment in Back-end

We use **Go**, **Gin**, **Mysql** and **Redis** for backend development and here is what we have done:

- Initialize the database

- Accomplish the connection to database API

- Add token parse and release function for JWT

- Create RESTful APIs --

- **Register**: for new user to create an account

- **Login**: to log in 

- **DeleteUser**: to delete an account

- **CreateCuisine**: to create new cuisines for a restaurant

- **DeleteCuisine**: to delete cuisines

- **ReadCuisine**: to get cuisines by restaurant

- **CreateOrder**: to create new orders for users

- **ReadOrder**: to get users' history orders

- **CreateRating**: to rate restaurants

- **GetRating**: to get a restaurant's rating

- **ReadRestaurant**: to get the restaurant list



## Environment

```
- MacBook Pro
- M1
- Visual Studio Code 2018
```
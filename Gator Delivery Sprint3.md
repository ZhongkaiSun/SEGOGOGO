# Gator Delivery Sprint 1



## Author Info

#### Back-end

Zhongkai Sun (sunzhongkai@ufl.edu 7888-5317)

Yudi Zheng (zhengyudi@ufl.edu 6715-2504)

#### Front-end

Yinghuan Zhang (zhangyinghuan@ufl.edu, 2741-2242) gitname: **Felixzhang666**

Hongru Liu (hongru.liu@ufl.edu, 5369-7439)



## Project Repository Link

https://github.com/ZhongkaiSun/SEGOGOGO

## Project Demo(Screenshots)
https://github.com/ZhongkaiSun/SEGOGOGO/blob/master/Project%20Demo(Screenshots).pdf

## Accomplishment in Front-end

For our Gator Delivery project, we implemented several pages using **HTML, CSS, JavaScript**, and **Bootstrap 5** in the front-end part:

- Home page: Base of our application, lead to other subpages. Includes functions like searching, paging and checkout. 

- History page: List previous orders and their details.

- Login page: Users can log in to accounts with a username and password.

- Sign-up page: Users can sign up for an account with their name, email, and create their password. We also add email validation and password checking here.

- Account information page: Users can review and change their personal information.

- Payment information page: Users can review and change their payment information. Add function to check login status using cookie.

- Change password page: Users can change their password if they forget it.

- Restaurant page: Show up the menu of restaurants.

- Restaurant cuisine page: Display restaurants of certain kind of cuisine.

  

The main workload is connected to the back-end, at stages 2-3.

-  Home page: Add the new shopping cart bottom, which can display the items that we want to order. Also, it has the checkout button to send the order messages to the server. Add the navigator of restaurant category to jump to different types of restaurants, like American, Asian, Coffee& Tea, etc.

- Login page: Using the ajax method to send user's information to the server and receive token message from the back-end. If it is successful, token will be stored in the cookie and show up an alert saying successful. Then go to the home page in logged-in status.

- Sign-up page: After users click the submit button, it will send the information to the server.

- Account information page: It will get the information from the server thought token, and it shows up at the place of the holder so that the user can modify it. Also, it can send information back to server to save it.

- Restaurant page: Get the restaurant basic information and menu from the server through the restaurant name.

- Restaurant cuisine page: Send the cuisine information to the server and get restaurant list of certain kind of cuisine.

- History page: Using the cookie's token to get and list previous orders and their details from the server.

  

Also, we use **cypress** to test each user interface. You can see it in the other files

##  Accomplishment in Back-end

We use **Go**, **Gin**, **Mysql,** and **Redis** for backend development, and here is what we have done:

- Initialize the database

- Accomplish the connection to database API

- Add token parse and release function for JWT

- Create RESTful APIs --

- **Register**: for new user to create an account

- **Login**: to log in 

- **DeleteUser**: to delete an account
  
- **ReadUser**: to get user's information

- **CreateCuisine**: to create new cuisines for a restaurant

- **DeleteCuisine**: to delete cuisines

- **ReadCuisine**: to get cuisines by restaurant

- **CreateOrder**: to create new orders for users

- **ReadOrder**: to get users' history orders

- **CreateRating**: to rate restaurants

- **GetRating**: to get a restaurant's rating

- **ReadRestaurant**: to get the restaurant list

- **CreatePayment**: to save a new payment method
  
- **ReadPayment**: get customers' payment information



## Environment

```
- MacBook Pro
- M1
- Visual Studio Code 2018
```
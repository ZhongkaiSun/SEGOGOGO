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



## Accomplishment in Front-end

For our Gator Delivery project, we implemented several pages using **HTML, CSS, JavaScript**, and **Bootstrap 5** in the front-end part:

- Home page: Base of our application, lead to another subpage. Includes functions like searching and paging.

- History page: List previous orders and their details.

- Login page: Users can log in to accounts with a username and password.

- Sign-up page: Users can sign-up for an account with their name, email, and create their password.

- Account information page: Users can review and change their personal information.

- Payment information page: Users can store and change their card information.

- Change password page: Users can change their password if they forget it.

- Restaurant page: Show up the information  of the restaurants

- Restaurant cusine page: Display the cosine of a restaurant.

  

The main workload is connected to the back-end, at stages 2-3.

-  Home page: Add the new shopping cart Botton, which can display the items that we want to order. Also, it has the checkout button to send the order messages to the server. Add the navigator of restaurant category to jump to the different types of restaurants like American, Asian, Coffee& Tea, etc.

- Login page: Using the ajax method sends the user's information to the server and receives the token message from the back-end. If it is successful, the cookie will store the token and show up an alert saying successful. Then go to the home page in logged-in status.

- Sign-up page: After users click the button of submitting, it will send the information to the server.

- Account information page: It will get the information from the server thought token, and it shows up at the place of the holder so that the user can modify it. Also, it can save to the server.

- Restaurant page: Get the restaurant information from the server through the restaurant name.

- Restaurant cusine page:  Get the cusine information from the server through the restaurant name.

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
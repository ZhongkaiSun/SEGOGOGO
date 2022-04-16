# Gator Delivery Sprint 4



## Author Info

#### Back-end

Zhongkai Sun (sunzhongkai@ufl.edu 7888-5317)

Yudi Zheng (zhengyudi@ufl.edu 6715-2504)

#### Front-end

Yinghuan Zhang (zhangyinghuan@ufl.edu, 2741-2242) gitname: **Felixzhang666**

Hongru Liu (hongru.liu@ufl.edu, 5369-7439)



## Project Repository Link

https://github.com/ZhongkaiSun/SEGOGOGO

## Project Demo(video)


## Accomplishment in Front-end

For our Gator Delivery project, we implemented several pages using **HTML, CSS, JavaScript**, and **Bootstrap 5** in the front-end part: 

At the before sprints, we have done the following page work and made them connect to the server successfully.



-  Home page: Add the new shopping cart bottom, which can display the items that we want to order. Also, it has the checkout button to send the order messages to the server. Add the navigator of restaurant category to jump to different types of restaurants, like American, Asian, Coffee& Tea, etc.

- Login page: Using the ajax method to send the user's information to the server and receive a token message from the back-end. If it is successful, the token will be stored in the cookie and show up an alert saying successful. Then go to the home page in logged-in status.

- Sign-up page: After users click the submit button, it will send the information to the server.

- Account information page: It will get the information from the server thought token, and it shows up at the place of the holder so that the user can modify it. Also, it can send information back to the server to save it.

- Restaurant page: Get the restaurant's basic information and menu from the server through the restaurant name.

- Restaurant cuisine page: Send the cuisine information to the server and get a restaurant list of certain kinds of cuisine.

- History page: Using the cookie's token to get and list previous orders and their details from the server.

  
  
  The main work of Sprint 4.
  
- Home page: Make a cart sidebar that can add and remove items and compute the total price. Modify the problem of removing items at the cart but cannot refresh the total price.

- Restaurant cuisine page: Make cart show up and check out the order. Make a cart sidebar that can add and remove items and compute the total price. 

- Restaurant page: Get the restaurant's Cuisine images from the server through the restaurant name. Make a cart sidebar that can add and remove items and compute the total price. 

- Login page: Initial more cookies to store cuisines and their prices.

- History page: Fix the problems of showing login or loginout Botton and alert when not at login status

- Account information page: Fix some show-up and send information bugs.  

  
  
  Also, we add or update test units of the home page,  restaurant, and restaurant cuisine page by **cypress**. You can see it in the other files

##  Accomplishment in Back-end

We use **Go**, **Gin**, **Mysql,** and **Redis** for backend development, and here is what we have done:





## Environment

```
- MacBook Pro
- M1
- Visual Studio Code 2018
```

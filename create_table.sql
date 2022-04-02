DROP TABLE IF EXISTS restaurants;

DROP TABLE IF EXISTS cuisines;

DROP TABLE IF EXISTS customers;

DROP TABLE IF EXISTS ratings;

DROP TABLE IF EXISTS ratings;

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS payments;

CREATE TABLE restaurants ( 
	id VARCHAR ( 50 ) PRIMARY KEY, 
	name VARCHAR ( 50 ), 
	address VARCHAR ( 100 ), 
	delivery_fee FLOAT, 
	img_path VARCHAR ( 100 ), 
	typeof_meal VARCHAR ( 50 ), 
	rating FLOAT 
);

CREATE TABLE cuisines ( 
	name VARCHAR ( 50 ), 
	restaurant_name VARCHAR ( 50 ), 
	price FLOAT, 
	calories INT,
	PRIMARY KEY (name, restaurant_name)
);

CREATE TABLE customers ( 
	username VARCHAR ( 50 ) PRIMARY KEY, 
	password VARCHAR ( 50 ) NOT NULL, 
	address_line1 VARCHAR ( 50 ) NOT NULL, 
	address_line2 VARCHAR ( 50 ), 
	phone VARCHAR ( 50 ),
	email VARCHAR ( 50 ) ,
	city VARCHAR ( 50 ),
	state VARCHAR ( 50 ),
	zipcode VARCHAR ( 50 ));


CREATE TABLE ratings ( 
	id int PRIMARY KEY AUTO_INCREMENT, 
	username VARCHAR (50),
	restaurant_id VARCHAR ( 50 ), 
	star INT, 
	rating_date date 
);

CREATE TABLE orders (
	username varchar(50),
	restaurant_name varchar(50),
	order_date varchar(50),
	price FLOAT,
	cuisine_name varchar(100),
	PRIMARY KEY(username, order_date, cuisine_name)
);

CREATE TABLE payments (
	username VARCHAR ( 50 ) PRIMARY KEY,
	card_holder VARCHAR ( 50 ),
 	card_number VARCHAR ( 50 ),
 	exp_date VARCHAR ( 50 ),
 	security_code VARCHAR ( 50 ),
 	address_line1 VARCHAR ( 50 ),
 	address_line2 VARCHAR ( 50 ),
 	city VARCHAR ( 50 ),
 	state VARCHAR ( 50 ),
 	zipcode VARCHAR ( 50 )
 );
-- dsd
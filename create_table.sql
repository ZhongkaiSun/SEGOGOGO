DROP TABLE IF EXISTS restaurants;

DROP TABLE IF EXISTS cuisines;

DROP TABLE IF EXISTS customers;

DROP TABLE IF EXISTS ratings;

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
	restaurant_id VARCHAR ( 50 ), 
	price FLOAT, 
	calories INT,
	PRIMARY KEY (name, restaurant_id)
);

CREATE TABLE customers ( 
	username VARCHAR ( 50 ) PRIMARY KEY, 
	password VARCHAR ( 50 ) NOT NULL, 
	address_line1 VARCHAR ( 50 ) NOT NULL, 
	address_line2 VARCHAR ( 50 ), 
	phone VARCHAR ( 50 ),
	email VARCHAR ( 50 ) );

CREATE TABLE ratings ( 
	id int PRIMARY KEY, 
	username VARCHAR (50),
	restaurant_id VARCHAR ( 50 ), 
	star INT, 
	rating_date date 
);

CREATE TABLE orders (
	username varchar(50) PRIMARY KEY,
	restaurantName varchar(50),
	orderDate varchar(50),
	price FLOAT,
	cuisineName varchar(100)
)

-- dsd
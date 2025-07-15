API Documentation
=================

This document provides comprehensive documentation for the Restaurant Management System API.

Table of Contents
----------------
1. Authentication
2. Users API
3. Food API
4. Menu API
5. Order API
6. Order Items API
7. Table API

1. Authentication
----------------
The API uses JWT (JSON Web Token) for authentication. Most endpoints require a valid token in the Authorization header.

To get a token:
1. Create an account using the signup endpoint
2. Login using your credentials
3. Use the received token in subsequent requests

2. Users API
-----------
Base URL: /users

Endpoints:

GET /users
- Description: Retrieve a list of users
- Authentication: Required
- Query Parameters:
  * recordPerPage (optional, default: 10)
  * page (optional, default: 1)
  * startIndex (optional)
- Response: Array of user objects

GET /users/:id
- Description: Retrieve a specific user by ID
- Authentication: Required
- Parameters:
  * id: User ID
- Response: User object

POST /users/signup
- Description: Create a new user account
- Authentication: Not required
- Request Body:
  {
    "first_name": "string",
    "last_name": "string",
    "email": "string",
    "password": "string",
    "phone": "string"
  }
- Response: Created user object with authentication tokens

POST /users/login
- Description: Authenticate user and get tokens
- Authentication: Not required
- Request Body:
  {
    "email": "string",
    "password": "string"
  }
- Response: User object with authentication tokens

3. Food API
----------
Base URL: /foods

Endpoints:

GET /foods
- Description: Retrieve a list of food items
- Authentication: Required
- Query Parameters:
  * recordPerPage (optional, default: 10)
  * page (optional, default: 1)
  * startIndex (optional)
- Response: Array of food items with pagination info

GET /foods/:food_id
- Description: Retrieve a specific food item by ID
- Authentication: Required
- Parameters:
  * food_id: Food item ID
- Response: Food item object

POST /foods
- Description: Create a new food item
- Authentication: Required
- Request Body:
  {
    "name": "string",
    "price": number,
    "food_image": "string",
    "menu_id": "string"
  }
- Response: Created food item object

PUT /foods/:id
- Description: Update a food item
- Authentication: Required
- Parameters:
  * id: Food item ID
- Request Body (all fields optional):
  {
    "name": "string",
    "price": number,
    "food_image": "string",
    "menu_id": "string"
  }
- Response: Update result object

4. Menu API
----------
Base URL: /menus

Endpoints:

GET /menus
- Description: Retrieve all menus
- Authentication: Required
- Response: Array of menu objects

GET /menus/:id
- Description: Retrieve a specific menu by ID
- Authentication: Required
- Parameters:
  * id: Menu ID
- Response: Menu object

POST /menus
- Description: Create a new menu
- Authentication: Required
- Request Body:
  {
    "name": "string",
    "category": "string",
    "start_date": "datetime",
    "end_date": "datetime"
  }
- Response: Created menu object

PUT /menus/:id
- Description: Update a menu
- Authentication: Required
- Parameters:
  * id: Menu ID
- Request Body (all fields optional):
  {
    "name": "string",
    "category": "string",
    "start_date": "datetime",
    "end_date": "datetime"
  }
- Response: Update result object
- Note: start_date must be after current time and end_date must be after start_date

5. Order API
----------
Base URL: /orders

Endpoints:

GET /orders
- Description: Retrieve all orders
- Authentication: Required
- Response: Array of order objects

GET /orders/:id
- Description: Retrieve a specific order by ID
- Authentication: Required
- Parameters:
  * id: Order ID
- Response: Order object

POST /orders
- Description: Create a new order
- Authentication: Required
- Request Body:
  {
    "table_id": "string" (optional)
  }
- Response: Created order object
- Note: If table_id is provided, it must exist in the database

PUT /orders/:id
- Description: Update an order
- Authentication: Required
- Parameters:
  * id: Order ID
- Request Body:
  {
    "table_id": "string" (optional)
  }
- Response: Update result object

6. Order Items API
---------------
Base URL: /orderItems

Endpoints:

GET /orderItems
- Description: Retrieve all order items
- Authentication: Required
- Response: Array of order item objects

GET /orderItems/:id
- Description: Retrieve a specific order item by ID
- Authentication: Required
- Parameters:
  * id: Order Item ID
- Response: Order item object

GET /orderItems/order/:id
- Description: Retrieve all order items for a specific order
- Authentication: Required
- Parameters:
  * id: Order ID
- Response: Detailed order items with food, order and table information

POST /orderItems
- Description: Create new order items
- Authentication: Required
- Request Body:
  {
    "table_id": "string",
    "order_items": [
      {
        "food_id": "string",
        "quantity": number,
        "unit_price": number
      }
    ]
  }
- Response: Created order items
- Note: Creates both the order and its items in one operation

PUT /orderItems/:id
- Description: Update an order item
- Authentication: Required
- Parameters:
  * id: Order Item ID
- Request Body (all fields optional):
  {
    "unit_price": number,
    "quantity": number,
    "food_id": "string"
  }
- Response: Update result object

7. Table API
----------
Base URL: /tables

Endpoints:

GET /tables
- Description: Retrieve all tables
- Authentication: Required
- Response: Array of table objects

GET /tables/:id
- Description: Retrieve a specific table by ID
- Authentication: Required
- Parameters:
  * id: Table ID
- Response: Table object

POST /tables
- Description: Create a new table
- Authentication: Required
- Request Body:
  {
    "table_number": number,
    "number_of_guests": number
  }
- Response: Created table object

PUT /tables/:id
- Description: Update a table
- Authentication: Required
- Parameters:
  * id: Table ID
- Request Body (all fields optional):
  {
    "table_number": number,
    "number_of_guests": number
  }
- Response: Update result object

Data Models
===========

1. User Model
------------
{
  "id": "ObjectId",
  "first_name": "string",
  "last_name": "string",
  "password": "string",
  "email": "string",
  "avatar": "string",
  "phone": "string",
  "token": "string",
  "refresh_token": "string",
  "created_at": "datetime",
  "updated_at": "datetime",
  "user_id": "string"
}

2. Food Model
------------
{
  "id": "ObjectId",
  "name": "string",
  "price": "number",
  "food_image": "string",
  "menu_id": "string",
  "created_at": "datetime",
  "updated_at": "datetime",
  "food_id": "string"
}

3. Menu Model
------------
{
  "id": "ObjectId",
  "name": "string",
  "category": "string",
  "start_date": "datetime",
  "end_date": "datetime",
  "created_at": "datetime",
  "updated_at": "datetime",
  "menu_id": "string"
}

4. Table Model
-------------
{
  "id": "ObjectId",
  "table_number": "number",
  "number_of_guests": "number",
  "created_at": "datetime",
  "updated_at": "datetime",
  "table_id": "string"
}

5. Order Model
-------------
{
  "id": "ObjectId",
  "table_id": "string",
  "created_at": "datetime",
  "updated_at": "datetime",
  "order_id": "string"
}

6. OrderItem Model
-----------------
{
  "id": "ObjectId",
  "quantity": "number",
  "unit_price": "number",
  "food_id": "string",
  "order_id": "string",
  "created_at": "datetime",
  "updated_at": "datetime",
  "order_item_id": "string"
}

API Documentation
===============

1. Authentication
----------------
The API uses JWT (JSON Web Token) for authentication. Most endpoints require a valid token in the Authorization header.

To get a token:
1. Create an account using the signup endpoint
2. Login using your credentials
3. Use the received token in subsequent requests

2. Users API
-----------
Base URL: /users

Endpoints:

GET /users
- Description: Retrieve a list of users
- Authentication: Required
- Query Parameters:
  * recordPerPage (optional, default: 10)
  * page (optional, default: 1)
  * startIndex (optional)
- Response: Array of user objects

GET /users/:id
- Description: Retrieve a specific user by ID
- Authentication: Required
- Parameters:
  * id: User ID
- Response: User object

POST /users/signup
- Description: Create a new user account
- Authentication: Not required
- Request Body:
  {
    "first_name": "string",
    "last_name": "string",
    "email": "string",
    "password": "string",
    "phone": "string"
  }
- Response: Created user object with authentication tokens

POST /users/login
- Description: Authenticate user and get tokens
- Authentication: Not required
- Request Body:
  {
    "email": "string",
    "password": "string"
  }
- Response: User object with authentication tokens

3. Food API
----------
Base URL: /foods

Endpoints:

GET /foods
- Description: Retrieve a list of food items
- Authentication: Required
- Query Parameters:
  * recordPerPage (optional, default: 10)
  * page (optional, default: 1)
  * startIndex (optional)
- Response: Array of food items with pagination info

GET /foods/:food_id
- Description: Retrieve a specific food item by ID
- Authentication: Required
- Parameters:
  * food_id: Food item ID
- Response: Food item object

POST /foods
- Description: Create a new food item
- Authentication: Required
- Request Body:
  {
    "name": "string",
    "price": number,
    "food_image": "string",
    "menu_id": "string"
  }
- Response: Created food item object

PUT /foods/:id
- Description: Update a food item
- Authentication: Required
- Parameters:
  * id: Food item ID
- Request Body (all fields optional):
  {
    "name": "string",
    "price": number,
    "food_image": "string",
    "menu_id": "string"
  }
- Response: Update result object

4. Menu API
----------
Base URL: /menus

Endpoints:

GET /menus
- Description: Retrieve all menus
- Authentication: Required
- Response: Array of menu objects

GET /menus/:id
- Description: Retrieve a specific menu by ID
- Authentication: Required
- Parameters:
  * id: Menu ID
- Response: Menu object

POST /menus
- Description: Create a new menu
- Authentication: Required
- Request Body:
  {
    "name": "string",
    "category": "string",
    "start_date": "datetime",
    "end_date": "datetime"
  }
- Response: Created menu object

PUT /menus/:id
- Description: Update a menu
- Authentication: Required
- Parameters:
  * id: Menu ID
- Request Body (all fields optional):
  {
    "name": "string",
    "category": "string",
    "start_date": "datetime",
    "end_date": "datetime"
  }
- Response: Update result object
- Note: start_date must be after current time and end_date must be after start_date

5. Order API
----------
Base URL: /orders

Endpoints:

GET /orders
- Description: Retrieve all orders
- Authentication: Required
- Response: Array of order objects

GET /orders/:id
- Description: Retrieve a specific order by ID
- Authentication: Required
- Parameters:
  * id: Order ID
- Response: Order object

POST /orders
- Description: Create a new order
- Authentication: Required
- Request Body:
  {
    "table_id": "string" (optional)
  }
- Response: Created order object
- Note: If table_id is provided, it must exist in the database

PUT /orders/:id
- Description: Update an order
- Authentication: Required
- Parameters:
  * id: Order ID
- Request Body:
  {
    "table_id": "string" (optional)
  }
- Response: Update result object

6. Order Items API
---------------
Base URL: /orderItems

Endpoints:

GET /orderItems
- Description: Retrieve all order items
- Authentication: Required
- Response: Array of order item objects

GET /orderItems/:id
- Description: Retrieve a specific order item by ID
- Authentication: Required
- Parameters:
  * id: Order Item ID
- Response: Order item object

GET /orderItems/order/:id
- Description: Retrieve all order items for a specific order
- Authentication: Required
- Parameters:
  * id: Order ID
- Response: Detailed order items with food, order and table information

POST /orderItems
- Description: Create new order items
- Authentication: Required
- Request Body:
  {
    "table_id": "string",
    "order_items": [
      {
        "food_id": "string",
        "quantity": number,
        "unit_price": number
      }
    ]
  }
- Response: Created order items
- Note: Creates both the order and its items in one operation

PUT /orderItems/:id
- Description: Update an order item
- Authentication: Required
- Parameters:
  * id: Order Item ID
- Request Body (all fields optional):
  {
    "unit_price": number,
    "quantity": number,
    "food_id": "string"
  }
- Response: Update result object

7. Table API
----------
Base URL: /tables

Endpoints:

GET /tables
- Description: Retrieve all tables
- Authentication: Required
- Response: Array of table objects

GET /tables/:id
- Description: Retrieve a specific table by ID
- Authentication: Required
- Parameters:
  * id: Table ID
- Response: Table object

POST /tables
- Description: Create a new table
- Authentication: Required
- Request Body:
  {
    "table_number": number,
    "number_of_guests": number
  }
- Response: Created table object

PUT /tables/:id
- Description: Update a table
- Authentication: Required
- Parameters:
  * id: Table ID
- Request Body (all fields optional):
  {
    "table_number": number,
    "number_of_guests": number
  }
- Response: Update result object
# User & Product Management Server

This project implements an HTTP web server using Go to manage users and products with different functionalities based on user roles. It uses PostgreSQL as the database and integrates Stripe for payment processing. The server exposes a set of REST APIs using the GorillaMux framework.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Server](#running-the-server)
- [API Endpoints](#api-endpoints)

## Features

### Normal User Endpoints
- **Sign Up User**
- **Login User**
- **Add Credit Card**
- **Delete Credit Card**
- **List Existing Products**
- **Buy (Multiple) Products**
- **Get User Bought Products History**

### Admin Endpoints
- **Create Product**
- **Update Product**
- **Delete Product**
- **Get Product Sales with Filtration**

## Technologies Used
- Go
- GorillaMux
- PostgreSQL
- Plain SQL (no ORM)

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/EbrahimAbdelaziz99/Rescounts_Task.git
    cd user-product-management
    ```

2. **Install Go dependencies:**
    ```bash
    go mod tidy
    ```

3. **Set up Docker:**
    Make sure you have Docker installed and running on your machine.

4. **Set up PostgreSQL database using Docker:**
    ```bash
    docker-compose up -d
    ```

5. **Create and initialize the database:**
    Execute the SQL scripts in `migrations/init.sql` to set up the database schema.

## Configuration

1. **Create a `.env` file in the root directory and add the following:**
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    STRIPE_SECRET_KEY=your_stripe_secret_key
    ```

## Running the Server

1. **Run the server:**
    ```bash
    go run main.go
    ```

2. The server will start on `http://localhost:8080`.

## API Endpoints

### Normal User Endpoints
- **POST /signup** - Sign up a new user
- **POST /login** - Login a user
- **POST /user/creditcard** - Add credit card
- **DELETE /user/creditcard** - Delete credit card
- **GET /products** - List existing products
- **POST /products/buy** - Buy (multiple) products
- **GET /user/products** - Get user bought products history

### Admin Endpoints
- **POST /admin/products** - Create product
- **PUT /admin/products/{id}** - Update product
- **DELETE /admin/products/{id}** - Delete product
- **GET /admin/products/sales** - Get product sales with filtration
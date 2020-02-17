# DDD Playground Golang

This is part of the ddd-playground project.

The project consists of many dockerised services written in different languages and exposing API Endpoints.

A gateway component written in Ballerina lang will expose transformed and combined data of those APIs.

## Content
The application is a product service exposing an basic API. 
It is written in Go.

### Get all products
```GET   /product```

### Get specific product
```GET   /product/{id}```

### Create a new product
```POST  /product```
```json
{
    "name": "Product Name",
    "categoryId": 1
}
```

## Initialize app

A Makefile is here to help.

Run `make` on your terminal and see what's available. 

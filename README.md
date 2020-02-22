# DDD Playground Golang

This is part of the ddd-playground project.

The project consists of many dockerised services written in different languages and exposing API Endpoints.

A gateway component written in Ballerina lang will expose transformed and combined data of those APIs.

## Content

The application is a product service exposing an basic API.
It is written in Go.

### Get all products

```json
curl --location --request GET 'http://localhost:8081/product'

[
   {
      "id":1,
      "name":"Red Chair",
      "category_id":1
   },
   {
      "id":2,
      "name":"Blue Table",
      "category_id":2
   },
   {
      "id":3,
      "name":"Yellow Carpet",
      "category_id":3
   }
]
```

### Get specific product

```json
curl --location --request GET 'http://localhost:8081/product/2'

{
   "id":2,
   "name":"Blue Table",
   "category_id":2
}
```

### Create a new product

```json
curl --location --request POST 'http://localhost:8081/product' -d '{"name":"Yellow Chair", "category_id":1}'
```

## Initialize app

Run `make init` to initialize your app.

## Import dummy data

Run `make install-fixtures` to import dummy data.

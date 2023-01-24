# survivor-game
![Generic badge](https://img.shields.io/badge/mantainer-isacarcanjo-black.svg)
![Generic badge](https://img.shields.io/badge/Golang-1.18.0-blue.svg)
![Generic badge](https://img.shields.io/badge/version-1.0.0-gree.svg)

## Description
Service responsible to stores survivors from a zombie apocalypse and allow
them to find the closest survivor from their location.

## Tech Stack
1. [Golang 1.18+](https://go.dev/blog/go1.18)

## Methods
1. Clean architecture
2. SOLID

## Routes
* /v1
  * [POST /v1/survivor](http://localhost:5000/v1/swagger/index.html#/survivor/post_survivor)
  
  * [GET /v1/survivor/closest-survivor/id](http://localhost:5000/v1/swagger/index.html#/survivor/get_survivor_closest_survivor__id_)

  * [GET /v1/survivor/id](http://localhost:5000/v1/swagger/index.html#/survivor/get_survivor__id_)
  
  * [PUT /v1/survivor/id](http://localhost:5000/v1/swagger/index.html#/survivor/put_survivor__id_)

  * [POST /v1/survivor/infected/id](http://localhost:5000/v1/swagger/index.html#/survivor/post_survivor_infected__id_)

## Local Variables

If not exist, you should make .env file
```
PORT=5000
VERSION=v1
SURVIVOR_COLLECTION=survivor
```

## Installing
The easier way is run this

```
go get ./...
``` 

or you can see the [documentation](https://go.dev/doc/modules/managing-dependencies)

## Run Application

How to run the application

```
go run main.go
```

## Documentation
http://localhost:5000/v1/swagger/index.html

## Requirements

You need to develop a RESTful API that stores survivors from a zombie apocalypse and allow
them to find the closest survivor from their location.

- ✅ Create new survivors - name, gender and current location (latitude/ longitude)
- ✅ Retrieve a survivor - survivor id required
- ✅Update a survivor
- ✅ Retrieve closest survivor from a survivor - survivor id required, you can use
only latitude or longitude to calculate it. This is to help a survivor to identify who is
closer.
- ✅ Mark survivor as infected - A survivor is marked infected when at least 3 others
survivors report that it is infected.
- ✅ Documentation + Setup - We need to run your code based on README and
use the API based on the documentation.

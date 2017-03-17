# Simple REST API
A simple REST API built in Golang. Based upon this screencast by Nic Raboy - https://www.youtube.com/watch?v=t96hBT53S4U

## Details
* No database attached here. Simple in memory mock data.
* Routing is handled by https://github.com/gorilla/mux

## Usage
Run the app like so:
```bash
$ go build && ./simple-rest-api
```

Then in your favourite REST client you can make requests to the following endpoints:
* Get all people - GET http://localhost:12345/people
* Get a person - GET http://localhost:12345/people/{id}
* Create a person - POST http://localhost:12345/people
* Delete a person - DELETE http://localhost:12345/people/{id}

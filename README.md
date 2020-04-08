### GO-SIMPLE-API

This is a simple API that I buildt in order to put in practice some Golang basic syntax and concepts.

#### HOW TO USE IT

- Clone this repo
- Move inside its directory: `cd go-simple-api`
- Compile the app: `go build main.go`
- Execute it: `./main`

##### ENDPOINTS

- GET all the persons: `http://localhost:3000/persons`
- GET one specific person: `http://localhost:3000/persons/{id}`
- POST (create) a new person: `http://localhost:3000/persons`. Here you have some data as example:
  ```
  {
      "firstname": "Bill",
      "lastname": "Murray",
      "job": "Genius"
  }
  ```
- PUT (update) one specific person: `http://localhost:3000/persons/{id}`
- DELETE one specific person: `http://localhost:3000/persons/{id}`

That's it. Be free to use it, change it and improve it.

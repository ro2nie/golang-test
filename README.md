## Name: Ronnie Winter

# Form3client

# Documentation of technical desicions
This can be found on the file named [TechnicalDecisions.md](TechnicalDecisions.md)

# Instructions

## Using library

Install the form3client to your `$GOBIN` path (usually `$GOPATH/bin`) by running the command `go install form3client`.
Then set the `$API` environment variable on your host to point to the address of form3's `accountapi` service (e.g: `http://localhost:8080/`).
Afterwards, you can create a new go project and use the library in this way: 
```
package main

import (
	form3client "github.com/form3interview"
)

func main() {
	// Create(). Param pointer to form3client.Account struct holding the account data
	form3client.Create(&form3client.Account{}) // Returns form3client.Data struct holding the created account, and error

	// Fetch(). Param id of account to fetch
	form3client.Fetch("<the id (uuid) of the account to fetch>") // Returns form3client.Data struct holding the fetched account, and error

	// List(). Params pageSize and pageNumber (optional)
	form3client.List(0, 20) // Returns form3client.DataList struct holding the fetched accounts, and error

	// Delete(). Params id of account to fetch and version
	form3client.Delete("<the id (uuid) of the account to fetch>", 0) // Returns status bool of whether or not the deletion was successful, and error.
}
```

## Running tests

In order to run tests locally, make sure you have set the environment variable `$API` to point to form3's `accountapi` (e.g: `http://localhost:8080/`)

You can also run tests as part of the docker-compose stack definition. For this, make sure you have Docker and docker-compose installed, then execute the command `docker-compose up`.

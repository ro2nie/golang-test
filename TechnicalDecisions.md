# Technical Decisions

## Separation of concern
The way I approached the design of this application was to have the functions `Create(), Fetch(), Delete(), List()` all within the same file (`client.go`).
These could in turn call the `rest.go` file functions which would deal with HTTP queries and errors. I kept all the marshalling separate, in `marshal.go`.
Each of the files only deals with one concern. This way the code is more testable and the units to test are simpler even though only bdd tests were required in this exercise.

## Variable parameter length
Coming from a Java background I thought that GO would offer method overloading which would let me call List() with or without parameters.
After I found that this was not possible, I resorted to using variadic parameters, which allowed me to pass none, 1, 2 or n number of parameters.
This created another problem. We now let the user pass an infinite number of int parameters. In order to solve this, I made it so that if the user was to pass
exactly 2 int parameters, the client will honour those as pageNumber and pageSize. Any other number of parameters (including 1 parameter) will return the same
as if the function List() was called without any parameters.

## docker-compose.yml file using bash script in entrypoint to wait for api
I had a few issues with the docker-compose stack definition, as the tests would not wait for the API to be active. The `depends_on` declarative helped in starting the 
client container last, but it would still fire too quickly, before the api was healthy. I tried using the following: 
```
healthcheck:
  test: [ "CMD", "curl", "-f", "http://accountapi/v1/health" ]
```
This did work sometimes, but it was still causing intermitent issues. 
In order to mitigate this, I resorted to writing my own bash script stored in `bin/start-tests.sh`, which loops through until the api becomes healthy. 

## Dockerfile used for creating the client container
I decided to use a Dockerfile that sets up a container, copies the app to it, sets the right environment variable to hit the `accountsapi` and installs the godog bdd dependency.
This way, from `docker-compose.yml` file, the `build: .` clause can build the container as part of the rest of the stack 

## Marshalling and unmarshalling into one Account struct
When I was using Postman to interact with the API, I noticed that the create account endpoint was replying with the same account but with added extra fields.
I decided to use the `json:"key,ommitempty"` declarative, in order to create an Account struct with data (which was missing the created_on, modified_on and version properties).
The ommitempty declarative helped me create a new Account that would have those fields as undefined. Later when the Create() function ran, the response would be unmarshalled
to contain these fields.

# Testing
For testing, I decided to use a BDD dependency that allowed me to write feature files and link them to real code.
In the tests I tried to cover any edge case I could by passing erroneous data where possible.
The objective of these tests were not to test the underlying API, but the actual client, to test every possible path, including: making the http requests,
marshalling and unmarshalling (of successful and error response), as well as error handling.
I wanted to make sure that the errors returned by the API would not be swallowed up, but that they would make it all the way up to the errors returned by the 4 functions.

# API validation errors repetition
I noticed that when sending invalid requests to the API, whenever a validation error message appeared, it returned something similar to this:
```
{
    "error_message": "validation failure list:\nvalidation failure list:\nvalidation failure list:\nbase_currency in body should match '^[A-Z]{3}$'\ncountry in body should match '^[A-Z]{2}$'"
}
```
I noticed that for all validation errors, the sentence `validation failure list:` was repeated three times. In order to return an error from the relevant operations, I decided
to sanitise this string, and remove the repetition. This is found in the `accounts.go` file on the `cleanErrorMessage()` function.

# Unmarshalling of Data and DataList
I tried to merge the `handleResponseData()` and `handleResponseDataList()` functions in `marshal.go` to work with both Data and DataList struct unmarshalling. My intention was 
to have the function return a dynamic type for code reusability. Unfortunately I was not able to find a simple way of solving this issue.
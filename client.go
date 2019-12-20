package form3client

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type request struct {
	method      string
	resource    string
	queryParams map[string]string
	payload     []byte
}

// Create will creates a new account by calling the accounts API.
// The function receives an account pointer of type Account
// which is the struct where all the information about the
// new account is stored.
//
// The function will return a Data struct which will contain the
// same data as the passed in account struct (dereferenced)
// with some extra properties such as createdOn, modifiedOn, version.
// The function also returns an error in the case that an error was encountered
func Create(account *Account) (Data, error) {
	payload, err := json.Marshal(Data{Account: *account})
	if err != nil {
		return Data{}, err
	}

	responseStatus, responsePayload, err := doRequest(&request{
		method:   "POST",
		resource: "v1/organisation/accounts/",
		payload:  payload,
	})

	if err != nil {
		return Data{}, err
	}

	return handleResponseData(responsePayload, http.StatusCreated, responseStatus)
}

// Fetch fetches an account specified by the account id
// The function takes in the id string of the id to retrieve.
// The function will return a Data struct which also contains an
// Account struct with all the information for that account
// The function also returns an error in the case that an error was encountered
func Fetch(id string) (Data, error) {
	responseStatus, responsePayload, err := doRequest(&request{
		method:   "GET",
		resource: "v1/organisation/accounts/" + id,
	})

	if err != nil {
		return Data{}, err
	}

	return handleResponseData(responsePayload, http.StatusOK, responseStatus)
}

// List will list a number of accounts stored by sending a request to the API.
// For pagination pass in page number and page size.
// The function takes integer varidic parameters, which specify the pageNumber and pageSize
// for paginating the account results. By accepting varidic integer parameters, the List() function
// allows for no pagination to be specified (when not passing any).
// If 2 integer parameters are passed, then the first will be taken for pageNumber and the second for pageSize.
// The function returns a DataList struct which contains a slice of Accounts returned.
// The function also returns an error in the case that an error was encountered
func List(params ...int) (DataList, error) {
	queryParams := map[string]string{}

	if len(params) == 2 {
		queryParams["page[number]"] = strconv.Itoa(params[0])
		queryParams["page[size]"] = strconv.Itoa(params[1])
	}

	responseStatus, responsePayload, err := doRequest(&request{
		method:      "GET",
		resource:    "v1/organisation/accounts/",
		queryParams: queryParams,
	})

	if err != nil {
		return DataList{}, err
	}
	return handleResponseDataList(responsePayload, http.StatusOK, responseStatus)
}

// Delete will delete an account with the given id and version
// The Delete() function takes in parameters for id to delete and the version number.
// The Delete() function will send a request to the API to delete an account record.
// The function wil return a boolean status of true for, 'successfully deleted' or false
// for 'not deleted'
// The function also returns an error in the case that an error was encountered
func Delete(id string, version int) (bool, error) {
	queryParams := map[string]string{
		"version": strconv.Itoa(version),
	}
	responseStatus, responsePayload, err := doRequest(&request{
		method:      "DELETE",
		resource:    "v1/organisation/accounts/" + id,
		queryParams: queryParams,
	})

	if err != nil {
		return false, err
	}

	return handleResponseDelete(responsePayload, http.StatusNoContent, responseStatus)
}

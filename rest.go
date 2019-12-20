package form3client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var commonHeaders = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
}
var client = &http.Client{}

func doRequest(req *request) (int, *[]byte, error) {
	request, err := http.NewRequest(req.method, os.Getenv("API")+req.resource, bytes.NewBuffer(req.payload))
	if err != nil {
		return 0, nil, fmt.Errorf("error: error whilst attempting to construct a %s request. %s", req.method, err)
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	addQueryParams(request, req.queryParams)

	response, err := client.Do(request)

	if err != nil {
		return 0, nil, fmt.Errorf("error: error whilst attempting to send a %s request. %s", req.method, err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("error: error whilst attempting to read the response from the %s request. %s", req.method, err)
	}
	return response.StatusCode, &data, nil
}

func addQueryParams(request *http.Request, params map[string]string) {
	query := request.URL.Query()

	if params != nil {
		for paramKey, paramValue := range params {
			query.Add(paramKey, paramValue)
		}
		request.URL.RawQuery = query.Encode()
	}
}

package form3client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleResponseDelete(responsePayload *[]byte, successHTTPResponseStatus int, responseStatus int) (bool, error) {
	if responseStatus == http.StatusNoContent {
		return true, nil
	}

	var newError GenericError
	err := json.Unmarshal(*responsePayload, &newError)

	if err != nil {
		return false, err
	}
	return false, fmt.Errorf("error: %s", newError.cleanErrorMessage())
}

func handleResponseData(responsePayload *[]byte, successHTTPResponseStatus int, responseStatus int) (Data, error) {
	if responseStatus == successHTTPResponseStatus {
		var newData Data
		err := json.Unmarshal(*responsePayload, &newData)
		if err != nil {
			return Data{}, err
		}
		return newData, nil
	}

	var newError GenericError
	err := json.Unmarshal(*responsePayload, &newError)
	if err != nil {
		return Data{}, err
	}
	return Data{}, fmt.Errorf("error: %s", newError.cleanErrorMessage())
}

func handleResponseDataList(responsePayload *[]byte, successHTTPResponseStatus int, responseStatus int) (DataList, error) {
	if responseStatus == http.StatusOK {
		var newData DataList
		err := json.Unmarshal(*responsePayload, &newData)
		if err != nil {
			return DataList{}, err
		}
		return newData, nil
	}

	var newError GenericError
	err := json.Unmarshal(*responsePayload, &newError)
	if err != nil {
		return DataList{}, err
	}
	return DataList{}, fmt.Errorf("error: %s", newError.cleanErrorMessage())
}

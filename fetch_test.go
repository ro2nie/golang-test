package form3client

import (
	"fmt"
	"reflect"
)

func (c *clientFeature) theFetchedDataShouldExactlyMatchTheCreatedData() error {
	if !reflect.DeepEqual(c.createdAccountResponse, c.fetchedAccount) {
		return fmt.Errorf("Expected fetched account response of: %+v\n to be equal to the Actual created account response: %+v\n but it was not", c.fetchedAccount, c.createdAccountResponse)
	}
	return nil
}

func (c *clientFeature) theUserUsesTheClientToFetchIt() error {
	data, err := Fetch(c.originalAccount.ID)
	c.fetchedAccount = data
	c.err = err
	return nil
}

package form3client

import (
	"fmt"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
)

func (c *clientFeature) theUserUsesTheClientToDeleteTheCreatedAccount() error {
	_, err := Delete(c.createdAccountResponse.ID, c.createdAccountResponse.Version)
	c.err = err
	return nil
}

func (c *clientFeature) theUserAttemptsToDeleteIt() error {
	status, err := Delete(c.originalAccount.ID, 0)
	c.err = err
	c.status = status
	return nil
}

func (c *clientFeature) theUserUsesTheClientToFetchTheCreatedAccount() error {
	data, err := Fetch(c.createdAccountResponse.ID)
	c.fetchedAccount = data
	c.err = err
	return nil
}

func (c *clientFeature) theUserShouldGetTheError(expectedErrorMessage *gherkin.DocString) error {
	expectedErrorMessageWithID := strings.Replace(expectedErrorMessage.Content, "<id>", c.createdAccountResponse.ID, 1)

	if c.err.Error() != expectedErrorMessageWithID {
		return fmt.Errorf("expected error message \n%s\n was not equal to the actual error message \n%s\n ", expectedErrorMessageWithID, c.err.Error())
	}
	return nil
}

func (c *clientFeature) theUserShouldGetASuccessfulResponse() error {
	if !c.status {
		return fmt.Errorf("expected response status to be \n%t\n but it was \n%t\n ", !c.status, c.status)
	}
	return nil
}

func (c *clientFeature) anAccountIDOf(id string) error {
	c.originalAccount.ID = id
	return nil
}

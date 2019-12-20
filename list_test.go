package form3client

import (
	"fmt"
)

func (c *clientFeature) theUserUsesTheClientToCreateNewAccountsBasedOnFile(numberOfAccounts int, filename string) error {
	n := 0
	c.originalAccount = readFromFile(filename)
	for n < numberOfAccounts {
		c.originalAccount.ID = generateUUID()
		c.createdAccountIDs = append(c.createdAccountIDs, c.originalAccount.ID)
		_, err := Create(&c.originalAccount)
		if err != nil {
			return fmt.Errorf("an error has occurred whilst attempting to create accounts")
		}
		n++
	}
	return nil
}

func (c *clientFeature) theUserListsTheAccountsWithDefaultPagination() error {
	dataList, err := List()
	c.err = err
	c.fetchedAccounts = dataList
	return nil
}

func (c *clientFeature) theUserShouldGetAResponseWithAccountsListed(numberOfAccounts int) error {
	if len(c.fetchedAccounts.Accounts) != numberOfAccounts {
		return fmt.Errorf("expected the number of fetched accounts to be %d but it was %d", numberOfAccounts, len(c.fetchedAccounts.Accounts))
	}
	return nil
}

func (c *clientFeature) theUserListsTheAccountsWithPageNumberPageSize(pageNumber, pageSize int) error {
	dataList, err := List(pageNumber, pageSize)
	c.err = err
	c.fetchedAccounts = dataList
	return nil
}

func (c *clientFeature) theUserListsTheAccountsWithPageNumber(pageNumber int) error {
	dataList, err := List(pageNumber)
	c.err = err
	c.fetchedAccounts = dataList
	return nil
}

func (c *clientFeature) theUserListsTheAccountsWithPageSize(pageSize int) error {
	//This is kind of pointless as client looks at array position 0 and 1 for page number and size.
	dataList, err := List(pageSize)
	c.err = err
	c.fetchedAccounts = dataList
	return nil
}

func (c *clientFeature) theUserListsTheAccountsWithPageNumberPageSizeExtraArgument(pageNumber, pageSize, extra int) error {
	dataList, err := List(pageNumber, pageSize, extra)
	c.err = err
	c.fetchedAccounts = dataList
	return nil
}

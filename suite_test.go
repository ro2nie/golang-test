package form3client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/DATA-DOG/godog"
)

type clientFeature struct {
	originalAccount        Account
	createdAccountResponse Data
	fetchedAccount         Data
	fetchedAccounts        DataList
	err                    error
	status                 bool
	createdAccountIDs      []string
}

func FeatureContext(s *godog.Suite) {
	originalHost := os.Getenv("API")
	client := &clientFeature{}

	s.Step(`^a new account is set up from file "([^"]*)"$`, client.ANewAccountIsSetUp)
	s.Step(`^the user uses the client to create it$`, client.theUserUsesTheClientToCreateIt)
	s.Step(`^the user should get a response with the same account plus extra fields$`, client.theUserShouldGetAResponseWithTheSameAccountPlusExtraFields)
	s.Step(`^the user should get the following error:$`, client.theUserShouldGetTheFollowingError)
	s.Step(`^the user should get an error containing:$`, client.theUserShouldGetAnErrorContaining)
	s.Step(`^the user has no network access$`, theUserHasNoNetworkAccess)
	s.Step(`^the user uses the client to delete the created account$`, client.theUserUsesTheClientToDeleteTheCreatedAccount)
	s.Step(`^the user uses the client to fetch the created account$`, client.theUserUsesTheClientToFetchTheCreatedAccount)
	s.Step(`^the user should get the error:$`, client.theUserShouldGetTheError)
	s.Step(`^an account ID of "([^"]*)"$`, client.anAccountIDOf)
	s.Step(`^the user attempts to delete it$`, client.theUserAttemptsToDeleteIt)
	s.Step(`^the user should get a successful response$`, client.theUserShouldGetASuccessfulResponse)
	s.Step(`^the fetched data should exactly match the created data$`, client.theFetchedDataShouldExactlyMatchTheCreatedData)
	s.Step(`^the user uses the client to fetch it$`, client.theUserUsesTheClientToFetchIt)
	s.Step(`^the user uses the client to create (\d+) new accounts based on file "([^"]*)"$`, client.theUserUsesTheClientToCreateNewAccountsBasedOnFile)
	s.Step(`^the user lists the accounts with default pagination$`, client.theUserListsTheAccountsWithDefaultPagination)
	s.Step(`^the user should get a response with (\d+) accounts listed$`, client.theUserShouldGetAResponseWithAccountsListed)
	s.Step(`^the user lists the accounts with page number (-?\d+) page size (-?\d+)$`, client.theUserListsTheAccountsWithPageNumberPageSize)
	s.Step(`^the user lists the accounts with page number (\d+)$`, client.theUserListsTheAccountsWithPageNumber)
	s.Step(`^the user lists the accounts with page size (\d+)$`, client.theUserListsTheAccountsWithPageSize)
	s.Step(`^the user lists the accounts with page number (\d+), page size (\d+), extra argument (\d+)$`, client.theUserListsTheAccountsWithPageNumberPageSizeExtraArgument)

	s.BeforeScenario(func(interface{}) {
		fmt.Println("Cleaning up test accounts")
		deleteAccounts(&client.createdAccountIDs)
		//Some tests will unset the env var to mimic a network error.
		client.fetchedAccount = Data{}
		client.createdAccountResponse = Data{}
		client.fetchedAccount = Data{}
		client.fetchedAccounts = DataList{}
		client.createdAccountIDs = []string{}
		client.err = nil
		client.status = false

		os.Setenv("API", originalHost)
	})

	s.AfterSuite(func() {
		fmt.Println("Cleaning up test accounts")
		deleteAccounts(&client.createdAccountIDs)
	})
}

func deleteAccounts(createdAccountIDs *[]string) {
	for _, id := range *createdAccountIDs {
		Delete(id, 0)
	}
}

func readFromFile(filename string) Account {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	var account Account
	err = json.Unmarshal(bytes, &account)
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
	return account
}

func generateUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

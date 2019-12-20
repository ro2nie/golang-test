package form3client

import (
	"strings"
	"time"
)

// DataList contains the returned accounts from issuing a List()
// It contains a struct that holds all returned accounts, as well
// as a struct of Links that holds links to the first and last records.
type DataList struct {
	Accounts []Account `json:"data"`
	Links    `json:"links"`
}

// Data contains the schema returned by Create() and Fetch().
// The Account holds the account data and the links hold
// links to the first and last records as well as itself
type Data struct {
	Account `json:"data"`
	Links   `json:"links"`
}

//Links contain links to the first and second account records as well as itself.
type Links struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Self  string `json:"self,omitempty"`
}

// Account is a struct that holds all the data with regards to a specific account
type Account struct {
	Attributes     `json:"attributes"`
	ID             string     `json:"id"`
	CreatedOn      *time.Time `json:"created_on,omitempty"`
	ModifiedOn     *time.Time `json:"modified_on,omitempty"`
	OrganisationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
	Version        int        `json:"version,omitempty"`
}

// Attributes are a sub struct that is a member of Account.
// It is used to hold some of the account's information.
type Attributes struct {
	AccountClassification       string   `json:"account_classification"`
	AccountMatchingOptOut       bool     `json:"account_matching_opt_out"`
	AccountNumber               string   `json:"account_number"`
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
	BankAccountName             string   `json:"bank_account_name"`
	BankID                      string   `json:"bank_id"`
	BankIDCode                  string   `json:"bank_id_code"`
	BaseCurrency                string   `json:"base_currency"`
	Bic                         string   `json:"bic"`
	Country                     string   `json:"country"`
	CustomerID                  string   `json:"customer_id"`
	FirstName                   string   `json:"first_name"`
	Iban                        string   `json:"iban"`
	JointAccount                bool     `json:"joint_account"`
	SecondaryIdentification     string   `json:"secondary_identification"`
	Title                       string   `json:"title"`
}

// GenericError is an error that is used to unmarshall errors
// thrown by the account api.
type GenericError struct {
	Message string `json:"error_message"`
}

// cleanErrorMessage() will remove 2 of the validation failure list messages.
// The API currently replies with validation failure list:\nvalidation failure list:\nvalidation failure list:\n
// which seems to be a excessive. This function makes the API validation error messages look a bit better.
func (e GenericError) cleanErrorMessage() string {
	return strings.Replace(e.Message, "validation failure list:\n", "", 2)
}

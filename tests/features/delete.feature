Feature: Account deletion

  Scenario: Create new account, then delete it and then attempt to fetch it
    Given a new account is set up from file "tests/resources/account.json"
    When the user uses the client to create it
    And the user uses the client to delete the created account
    And the user uses the client to fetch the created account
    Then the user should get the error:
      """
      error: record <id> does not exist
      """

  Scenario: Delete an nonexistent account
    Given an account ID of "498be16f-828f-497d-b1ff-e6f7eb9b0ec1"
    When the user attempts to delete it
    Then the user should get a successful response

  Scenario: Delete an incorrect account id
    Given an account ID of "incorrect-account-id"
    When the user attempts to delete it
    Then the user should get the error:
      """
      error: id is not a valid uuid
      """

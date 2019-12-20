Feature: Account deletion

  Scenario: Create 30 accounts and default list"
    Given the user uses the client to create 50 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with default pagination
    Then the user should get a response with 50 accounts listed

  Scenario: Create 30 accounts and paginate list page 0, size 30"
    Given the user uses the client to create 30 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number 0 page size 30
    Then the user should get a response with 30 accounts listed

  Scenario: Create 20 accounts and paginate list page 0, size 15"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number 0 page size 15
    Then the user should get a response with 15 accounts listed

  Scenario: Create 20 accounts and paginate list page 1, size 5"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number 1 page size 5
    Then the user should get a response with 5 accounts listed

  Scenario: Create 20 accounts and paginate list page 0, no size"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number 0
    Then the user should get a response with 20 accounts listed

  Scenario: Create 20 accounts and paginate list page 100, no size"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number 100
    Then the user should get a response with 20 accounts listed

  Scenario: Create 20 accounts and paginate list no page, size 5"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page size 1
    Then the user should get a response with 20 accounts listed

  Scenario: Create 20 accounts and paginate list no page, size 5"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number 0, page size 1, extra argument 100
    Then the user should get a response with 20 accounts listed

  Scenario: Create 20 accounts and paginate list page 1, size 40"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number 1 page size 40
    Then the user should get a response with 0 accounts listed

  Scenario: Create 20 accounts and paginate list page -1, size -20"
    Given the user uses the client to create 20 new accounts based on file "tests/resources/account.json"
    When the user lists the accounts with page number -1 page size -20
    Then the user should get a response with 0 accounts listed
    And the user should get the following error:
      """
      error: server error
      """
# Email Validator Library

## Overview

The Email Validator Library is a robust and flexible Go library designed to validate email addresses using a chain of validators. This library includes various checks such as syntax validation, role-based email checks, MX records verification, SMTP validation, and more.

## Features

- Syntax Validation
- Role-Based Email Check
- MX Record Verification
- SMTP Validation
- Ban Words in Username Check
- Blacklisted Emails Check
- DNS Check
- RFC Compliance Check
- No RFC Warnings Check
- Gravatar Check

## Installation

To install the library, use the following command:

```sh
go get github.com/sufimalek/emailvalidator
```

## Project Layout ##
```
emailvalidator/
|-- main.go
|-- validators/
    |-- ban_words_username_validator.go
    |-- blacklist_emails_validator.go
    |-- dns_check_validator.go
    |-- gravatar_validator.go
    |-- mx_validator.go
    |-- no_rfc_warnings_validator.go
    |-- rfc_validator.go
    |-- role_validator.go
    |-- smtp_validator.go
    |-- syntax_validator.go
|-- chain.go
|-- chain_test.go
```

## Usage

Here is an example of how to use the Email Validator Library:

```go
package main

import (
    "fmt"
    "github.com/sufimalek/emailvalidator"
    "github.com/sufimalek/emailvalidator/validators"
)

func main() {
    chain := emailvalidator.NewValidatorChainBuilder().
        Add(&validators.SyntaxValidator{}).
        Add(&validators.RoleValidator{}).
        Add(&validators.MXValidator{}).
        Add(&validators.SMTPValidator{}).
        Add(validators.NewBanWordsUsernameValidator([]string{"spam", "junk"})).
        Add(validators.NewBlackListEmailsValidator([]string{"blacklisted@example.com"})).
        Add(&validators.DNSCheckValidator{}).
        Add(&validators.RFCValidator{}).
        Add(&validators.NoRFCWarningsValidator{}).
        Add(&validators.GravatarValidator{}).
        Build()

    email := "user@example.com"
    if chain.Validate(email) {
        fmt.Printf("%s is a valid email!\n", email)
    } else {
        fmt.Printf("%s is not a valid email.\n", email)
    }
}
```

## Validators

Each validator is implemented as a separate component. Here are the available validators:

- `SyntaxValidator`: Checks the basic syntax of the email.
- `RoleValidator`: Ensures the email is not a role-based email.
- `MXValidator`: Verifies the domain has MX records.
- `SMTPValidator`: Attempts to connect to the mail server to verify the email.
- `BanWordsUsernameValidator`: Ensures the username does not contain banned words.
- `BlackListEmailsValidator`: Checks the email against a blacklist.
- `DNSCheckValidator`: Verifies the domain has valid DNS records.
- `RFCValidator`: Checks if the email adheres to RFC standards.
- `NoRFCWarningsValidator`: Ensures there are no RFC warnings.
- `GravatarValidator`: Checks if the email has a Gravatar.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
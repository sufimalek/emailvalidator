# Email validator package

## What is this repository for? ##

This is email validator repo


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

## 1. Add to project

Launch in the your project directory

```sh
go get github.com/sufimalek/emailvalidator
```


## 2. Usage validator
```go

package main

import (
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
		// Add(&validators.GravatarValidator{}).
		Build()

	isValid := chain.Validate(req.Email)

}
```
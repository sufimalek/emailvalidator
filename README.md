Project Structure

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
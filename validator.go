package emailvalidator

type Validator interface {
	Validate(email string) bool
}

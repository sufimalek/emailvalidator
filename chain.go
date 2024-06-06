package emailvalidator

type ValidatorChain struct {
	validators []Validator
}

func (vc *ValidatorChain) Validate(email string) bool {
	for _, validator := range vc.validators {
		if !validator.Validate(email) {
			return false
		}
	}
	return true
}

type ValidatorChainBuilder struct {
	validators []Validator
}

func NewValidatorChainBuilder() *ValidatorChainBuilder {
	return &ValidatorChainBuilder{}
}

func (b *ValidatorChainBuilder) Add(validator Validator) *ValidatorChainBuilder {
	b.validators = append(b.validators, validator)
	return b
}

func (b *ValidatorChainBuilder) Build() *ValidatorChain {
	return &ValidatorChain{validators: b.validators}
}

// package validators

// import "strings"

// type BlackListEmailsValidator struct {
// 	blacklistedEmails []string
// }

// func NewBlackListEmailsValidator(blacklistedEmails []string) *BlackListEmailsValidator {
// 	return &BlackListEmailsValidator{blacklistedEmails: blacklistedEmails}
// }

// func (v *BlackListEmailsValidator) Validate(email string) bool {
// 	for _, blacklisted := range v.blacklistedEmails {
// 		if strings.EqualFold(email, blacklisted) {
// 			return false
// 		}
// 	}
// 	return true
// }

package validators

import (
	"log"
)

type BlackListEmailsValidator struct {
	blacklistedEmails map[string]bool
}

func NewBlackListEmailsValidator(emails []string) *BlackListEmailsValidator {
	blacklist := make(map[string]bool)
	for _, email := range emails {
		blacklist[email] = true
	}
	return &BlackListEmailsValidator{blacklistedEmails: blacklist}
}

func (v *BlackListEmailsValidator) Validate(email string) bool {
	if v.blacklistedEmails[email] {
		log.Printf("BlackListEmailsValidator: Email %s is blacklisted", email)
		return false
	}
	return true
}

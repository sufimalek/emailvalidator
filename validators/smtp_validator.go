package validators

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

type SMTPValidator struct{}

func (v *SMTPValidator) Validate(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		log.Printf("SMTPValidator: Invalid email format for %s", email)
		return false
	}
	domain := parts[1]

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		log.Printf("SMTPValidator: No MX records found for domain %s", domain)
		return false
	}

	client, err := smtp.Dial(fmt.Sprintf("%s:25", mxRecords[0].Host))
	if err != nil {
		log.Printf("SMTPValidator: Could not connect to SMTP server for domain %s", domain)
		return false
	}
	defer client.Close()

	err = client.Mail("validator@example.com")
	if err != nil {
		log.Printf("SMTPValidator: Could not set sender for %s", domain)
		return false
	}

	err = client.Rcpt(email)
	if err != nil {
		log.Printf("SMTPValidator: Recipient email %s is not deliverable", email)
		return false
	}

	return true
}

// func (v *SMTPValidator) Validate(email string) bool {
// 	domain := strings.Split(email, "@")[1]
// 	mxRecords, err := net.LookupMX(domain)
// 	if err != nil || len(mxRecords) == 0 {
// 		return false
// 	}

// 	client, err := smtp.Dial(mxRecords[0].Host + ":25")
// 	if err != nil {
// 		return false
// 	}
// 	defer client.Close()
// 	return true
// }

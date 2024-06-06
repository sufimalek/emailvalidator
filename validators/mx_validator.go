package validators

import (
	"log"
	"net"
	"strings"
)

type MXValidator struct{}

func (v *MXValidator) Validate(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		log.Printf("MXValidator: Invalid email format for %s", email)
		return false
	}
	domain := parts[1]
	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		log.Printf("MXValidator: No MX records found for domain %s", domain)
		return false
	}
	return true
}

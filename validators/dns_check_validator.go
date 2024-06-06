package validators

import (
	"log"
	"net"
	"strings"
)

type DNSCheckValidator struct{}

func (v *DNSCheckValidator) Validate(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		log.Printf("DNSCheckValidator: Invalid email format for %s", email)
		return false
	}
	domain := parts[1]
	_, err := net.LookupHost(domain)
	if err != nil {
		log.Printf("DNSCheckValidator: No DNS records found for domain %s", domain)
		return false
	}
	return true
}

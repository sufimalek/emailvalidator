package validators

import (
	"log"
	"strings"
)

type RoleValidator struct{}

var roleEmails = []string{
	"admin", "administrator", "webmaster", "hostmaster", "postmaster",
	"info", "support", "sales",
}

func (v *RoleValidator) Validate(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		log.Printf("RoleValidator: Invalid email format for %s", email)
		return false
	}
	localPart := parts[0]
	for _, role := range roleEmails {
		if localPart == role {
			log.Printf("RoleValidator: Role-based email %s is not allowed", email)
			return false
		}
	}
	return true
}

// func (v *RoleValidator) Validate(email string) bool {
// 	roleEmails := []string{"admin", "support", "info"}
// 	for _, role := range roleEmails {
// 		if strings.HasPrefix(email, role) {
// 			log.Printf("RoleValidator: %s is a role-based email", email)
// 			return false
// 		}
// 	}
// 	log.Printf("RoleValidator: %s passed", email)
// 	return true
// }

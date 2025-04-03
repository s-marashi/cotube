package err

import "fmt"

type DomainError struct {
	domain  string
	entity  string
	field   string
	message string
}

func NewDomainError(domain, entity, field, message string) DomainError {
	return DomainError{
		domain:  domain,
		entity:  entity,
		field:   field,
		message: message,
	}
}

func (e DomainError) Error() string {
	return fmt.Sprintf("[%s] in entity %q field %q has invalid value: %s", e.domain, e.entity, e.field, e.message)
}

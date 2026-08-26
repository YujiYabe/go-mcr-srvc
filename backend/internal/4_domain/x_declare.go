package domain

import groupObject "backend/internal/4_domain/group_object"

// NewDomain ...
func NewDomain() ToDomain {
	return &domain{}
}

type (
	domain struct{}

	// ToDomain ...
	ToDomain interface {
		EnsurePrimaryEmploymentAssignable(
			user groupObject.User,
			userEmployment groupObject.UserEmployment,
		) error
	}
)

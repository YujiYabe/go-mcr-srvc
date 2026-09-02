package domain

import groupObject "backend/internal/4_domain/group_object"

// NewDomain ...
func NewDomain() (
	toDomain ToDomain,
) {
	toDomain = &domain{}
	return
}

type (
	domain struct{}

	// ToDomain ...
	ToDomain interface {
		EnsurePrimaryEmploymentAssignable(
			user groupObject.User,
			userEmployment groupObject.UserEmployment,
		) (
			err error,
		)
	}
)

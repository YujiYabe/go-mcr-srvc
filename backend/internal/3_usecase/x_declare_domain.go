package usecase

import groupObject "backend/internal/4_domain/group_object"

type (
	// ToDomain ...
	ToDomain interface {
		EnsurePrimaryEmploymentAssignable(
			user groupObject.User,
			userEmployment groupObject.UserEmployment,
		) error
	}
)

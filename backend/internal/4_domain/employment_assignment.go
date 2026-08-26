package domain

import (
	groupObject "backend/internal/4_domain/group_object"
	serviceObject "backend/internal/4_domain/service_object"
)

func (receiver *domain) EnsurePrimaryEmploymentAssignable(
	user groupObject.User,
	userEmployment groupObject.UserEmployment,
) error {
	return serviceObject.NewPrimaryEmploymentAssignmentPolicy().
		EnsureAssignable(user, userEmployment)
}

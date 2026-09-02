package domain

import (
	groupObject "backend/internal/4_domain/group_object"
	serviceObject "backend/internal/4_domain/service_object"
)

func (receiver *domain) EnsurePrimaryEmploymentAssignable(
	user groupObject.User,
	userEmployment groupObject.UserEmployment,
) (
	err error,
) {
	err = serviceObject.NewPrimaryEmploymentAssignmentPolicy().
		EnsureAssignable(user, userEmployment)
	return
}

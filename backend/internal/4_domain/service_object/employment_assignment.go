package service_object

import (
	"fmt"

	groupObject "backend/internal/4_domain/group_object"
)

type PrimaryEmploymentAssignmentPolicy struct{}

func NewPrimaryEmploymentAssignmentPolicy() (
	primaryEmploymentAssignmentPolicy *PrimaryEmploymentAssignmentPolicy,
) {
	primaryEmploymentAssignmentPolicy = &PrimaryEmploymentAssignmentPolicy{}

	return
}

func (receiver *PrimaryEmploymentAssignmentPolicy) EnsureAssignable(
	user groupObject.User,
	userEmployment groupObject.UserEmployment,
) (
	err error,
) {
	if !user.HasIdentity() {
		err = fmt.Errorf("user identity is required")

		return
	}
	if returnedErr := userEmployment.EnsureReadyToAssign(); returnedErr != nil {
		err = returnedErr

		return
	}
	if userEmployment.UserID().GetValue() != user.Identity().GetValue() {
		err = fmt.Errorf("primary employment must belong to the user")

		return
	}
	if !userEmployment.IsPrimary() {
		err = fmt.Errorf("primary employment assignment requires primary employment")

		return
	}

	err = nil

	return
}

package service_object

import (
	"fmt"

	groupObject "backend/internal/4_domain/group_object"
)

type PrimaryEmploymentAssignmentPolicy struct{}

func NewPrimaryEmploymentAssignmentPolicy() *PrimaryEmploymentAssignmentPolicy {
	return &PrimaryEmploymentAssignmentPolicy{}
}

func (receiver *PrimaryEmploymentAssignmentPolicy) EnsureAssignable(
	user groupObject.User,
	userEmployment groupObject.UserEmployment,
) error {
	if !user.HasIdentity() {
		return fmt.Errorf("user identity is required")
	}
	if err := userEmployment.EnsureReadyToAssign(); err != nil {
		return err
	}
	if userEmployment.UserID().GetValue() != user.Identity().GetValue() {
		return fmt.Errorf("primary employment must belong to the user")
	}
	if !userEmployment.IsPrimary() {
		return fmt.Errorf("primary employment assignment requires primary employment")
	}

	return nil
}

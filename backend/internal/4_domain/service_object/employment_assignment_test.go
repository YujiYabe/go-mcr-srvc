package service_object

import (
	"strings"
	"testing"

	groupObject "backend/internal/4_domain/group_object"
)

func TestPrimaryEmploymentAssignmentPolicyAllowsPrimaryEmployment(
	t *testing.T,
) {
	policy := NewPrimaryEmploymentAssignmentPolicy()
	user := newTestUser(t, 1)
	employment := newTestUserEmployment(t, 1, true)

	if err := policy.EnsureAssignable(user, employment); err != nil {
		t.Fatalf("expected assignment success, got: %v", err)
	}
}

func TestPrimaryEmploymentAssignmentPolicyRejectsDifferentUser(
	t *testing.T,
) {
	policy := NewPrimaryEmploymentAssignmentPolicy()
	user := newTestUser(t, 1)
	employment := newTestUserEmployment(t, 2, true)

	err := policy.EnsureAssignable(user, employment)
	if err == nil {
		t.Fatal("expected ownership error")
	}
	if !strings.Contains(err.Error(), "must belong to the user") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPrimaryEmploymentAssignmentPolicyRejectsNonPrimaryEmployment(
	t *testing.T,
) {
	policy := NewPrimaryEmploymentAssignmentPolicy()
	user := newTestUser(t, 1)
	employment := newTestUserEmployment(t, 1, false)

	err := policy.EnsureAssignable(user, employment)
	if err == nil {
		t.Fatal("expected primary employment error")
	}
	if !strings.Contains(err.Error(), "requires primary employment") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func newTestUser(
	t *testing.T,
	id int,
) (
	user groupObject.User,
) {
	t.Helper()

	name := "alice"
	email := "alice@example.com"
	reconstructedUser, err := groupObject.ReconstructUser(&groupObject.NewUserArgs{
		ID:    &id,
		Name:  &name,
		Email: &email,
	})
	if err != nil {
		t.Fatalf("failed to reconstruct user: %v", err)
	}

	user = *reconstructedUser

	return
}

func newTestUserEmployment(
	t *testing.T,
	userID int,
	isPrimary bool,
) (
	userEmployment groupObject.UserEmployment,
) {
	t.Helper()

	companyID := 1
	departmentID := 2
	positionID := 3
	employeeCode := "EMP001"
	employmentType := "full_time"
	newUserEmployment, err := groupObject.NewUserEmployment(&groupObject.NewUserEmploymentArgs{
		UserID:         &userID,
		CompanyID:      &companyID,
		DepartmentID:   &departmentID,
		PositionID:     &positionID,
		EmployeeCode:   &employeeCode,
		EmploymentType: &employmentType,
		IsPrimary:      &isPrimary,
	})
	if err != nil {
		t.Fatalf("failed to create user employment: %v", err)
	}

	userEmployment = *newUserEmployment

	return
}

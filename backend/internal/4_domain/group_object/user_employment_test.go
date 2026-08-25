package group_object

import (
	"strings"
	"testing"
)

func TestUserEmploymentRequiresAssignmentIdentities(t *testing.T) {
	employment, err := NewUserEmployment(&NewUserEmploymentArgs{
		UserID:       intPointer(1),
		DepartmentID: intPointer(1),
		PositionID:   intPointer(1),
	})
	if err != nil {
		t.Fatalf("failed to create user employment: %v", err)
	}

	err = employment.EnsureReadyToAssign()
	if err == nil {
		t.Fatal("expected company identity error")
	}
	if !strings.Contains(err.Error(), "company identity is required") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUserEmploymentCanValidateUserOwnership(t *testing.T) {
	user, err := ReconstructUser(&NewUserArgs{
		ID:    intPointer(1),
		Name:  stringPointer("alice"),
		Email: stringPointer("alice@example.com"),
	})
	if err != nil {
		t.Fatalf("failed to reconstruct user: %v", err)
	}
	employment := newTestUserEmployment(t, 1)

	if err := employment.EnsureBelongsTo(*user); err != nil {
		t.Fatalf("expected ownership success, got: %v", err)
	}
}

func TestUserEmploymentRejectsDifferentUser(t *testing.T) {
	user, err := ReconstructUser(&NewUserArgs{
		ID:    intPointer(1),
		Name:  stringPointer("alice"),
		Email: stringPointer("alice@example.com"),
	})
	if err != nil {
		t.Fatalf("failed to reconstruct user: %v", err)
	}
	employment := newTestUserEmployment(t, 2)

	err = employment.EnsureBelongsTo(*user)
	if err == nil {
		t.Fatal("expected ownership error")
	}
	if !strings.Contains(err.Error(), "must belong to the user") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func newTestUserEmployment(
	t *testing.T,
	userID int,
) UserEmployment {
	t.Helper()

	employeeCode := "EMP001"
	employmentType := "full_time"
	isPrimary := true
	employment, err := NewUserEmployment(&NewUserEmploymentArgs{
		UserID:         intPointer(userID),
		CompanyID:      intPointer(1),
		DepartmentID:   intPointer(2),
		PositionID:     intPointer(3),
		EmployeeCode:   &employeeCode,
		EmploymentType: &employmentType,
		IsPrimary:      &isPrimary,
	})
	if err != nil {
		t.Fatalf("failed to create user employment: %v", err)
	}

	return *employment
}

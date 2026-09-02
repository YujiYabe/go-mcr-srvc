package group_object

import (
	"strings"
	"testing"
)

func TestUserEmploymentRequiresAssignmentIdentities(
	t *testing.T,
) {
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

func TestUserEmploymentExposesAssignmentValues(
	t *testing.T,
) {
	employment := newTestUserEmployment(t, 1)

	if employment.UserID().GetValue() != 1 {
		t.Fatalf("expected user id 1, got: %d", employment.UserID().GetValue())
	}
	if !employment.IsPrimary() {
		t.Fatal("expected primary employment")
	}
	if employment.EmployeeCode() != "EMP001" {
		t.Fatalf("expected employee code EMP001, got: %s", employment.EmployeeCode())
	}
}

func newTestUserEmployment(
	t *testing.T,
	userID int,
) (
	value UserEmployment,
) {
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

	value = *employment

	return
}

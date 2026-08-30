package type_object

import "testing"

func TestNewPermissionListReturnsValuesInOrder(
	t *testing.T,
) {
	t.Parallel()

	values := []string{"user:read", "user:write"}

	permissionList, err := NewPermissionList(values)
	if err != nil {
		t.Fatalf("expected valid permission list, got error: %v", err)
	}

	actual := permissionList.GetSliceValue()
	if len(actual) != len(values) {
		t.Fatalf("expected %d values, got %d", len(values), len(actual))
	}
	for index := range values {
		if actual[index] != values[index] {
			t.Fatalf("expected value at %d to be %q, got %q", index, values[index], actual[index])
		}
	}
}

func TestNewPermissionListRejectsInvalidPermission(
	t *testing.T,
) {
	t.Parallel()

	values := []string{"user:read", ""}

	_, err := NewPermissionList(values)
	if err == nil {
		t.Fatal("expected invalid permission list to return an error")
	}
}

func TestPermissionListCanJudgeUserPermissions(
	t *testing.T,
) {
	t.Parallel()

	permissionList, err := NewPermissionList([]string{
		PermissionUserRead,
	})
	if err != nil {
		t.Fatalf("expected valid permission list, got error: %v", err)
	}

	if permissionList.IsEmpty() {
		t.Fatal("expected permission list not to be empty")
	}
	if permissionList.Count() != 1 {
		t.Fatalf("expected permission count 1, got %d", permissionList.Count())
	}
	if !permissionList.CanReadUser() {
		t.Fatal("expected user read permission")
	}
	if permissionList.CanWriteUser() {
		t.Fatal("did not expect user write permission")
	}
	if err := permissionList.EnsureHas(PermissionUserRead); err != nil {
		t.Fatalf("expected read permission, got error: %v", err)
	}
	if err := permissionList.EnsureHas(PermissionUserWrite); err == nil {
		t.Fatal("expected missing write permission error")
	}
}

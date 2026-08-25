package type_object

import "testing"

func TestNewPermissionListReturnsValuesInOrder(t *testing.T) {
	t.Parallel()

	values := []string{"person:read", "person:write"}

	permissionList, err := NewPermissionList(values)
	if err != nil {
		t.Fatalf("expected valid permission list, got error: %v", err)
	}

	actual := permissionList.GetSliceValue()
	if len(actual) != len(values) {
		t.Fatalf("expected %d values, got %d", len(values), len(actual))
	}
	for i := range values {
		if actual[i] != values[i] {
			t.Fatalf("expected value at %d to be %q, got %q", i, values[i], actual[i])
		}
	}
}

func TestNewPermissionListRejectsInvalidPermission(t *testing.T) {
	t.Parallel()

	values := []string{"person:read", ""}

	_, err := NewPermissionList(values)
	if err == nil {
		t.Fatal("expected invalid permission list to return an error")
	}
}

func TestPermissionListCanJudgePersonPermissions(t *testing.T) {
	t.Parallel()

	permissionList, err := NewPermissionList([]string{
		PermissionPersonRead,
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
	if !permissionList.CanReadPerson() {
		t.Fatal("expected person read permission")
	}
	if permissionList.CanWritePerson() {
		t.Fatal("did not expect person write permission")
	}
	if err := permissionList.EnsureHas(PermissionPersonRead); err != nil {
		t.Fatalf("expected read permission, got error: %v", err)
	}
	if err := permissionList.EnsureHas(PermissionPersonWrite); err == nil {
		t.Fatal("expected missing write permission error")
	}
}

package group_object

import "testing"

func TestNewPersonAllowsDraftWithoutIdentity(t *testing.T) {
	person, err := NewPerson(&NewPersonArgs{
		Name:        stringPointer("alice"),
		MailAddress: stringPointer("alice@example.com"),
	})
	if err != nil {
		t.Fatalf("expected draft person, got: %v", err)
	}

	if person.HasIdentity() {
		t.Fatal("draft person should not have identity")
	}
	if !person.CanBeUsedAsSearchCondition() {
		t.Fatal("person with name should be usable as search condition")
	}
}

func TestReconstructPersonRequiresIdentity(t *testing.T) {
	_, err := ReconstructPerson(&NewPersonArgs{
		Name:        stringPointer("alice"),
		MailAddress: stringPointer("alice@example.com"),
	})
	if err == nil {
		t.Fatal("expected identity error")
	}
}

func TestNewPersonSearchConditionRequiresCondition(t *testing.T) {
	_, err := NewPersonSearchCondition(&NewPersonArgs{})
	if err == nil {
		t.Fatal("expected search condition error")
	}
}

func TestPersonCanChangeNameAndMailAddress(t *testing.T) {
	person, err := ReconstructPerson(&NewPersonArgs{
		ID:          intPointer(1),
		Name:        stringPointer("alice"),
		MailAddress: stringPointer("alice@example.com"),
	})
	if err != nil {
		t.Fatalf("failed to reconstruct person: %v", err)
	}

	if err := person.Rename(stringPointer("bob")); err != nil {
		t.Fatalf("expected rename success, got: %v", err)
	}
	if err := person.ChangeMailAddress(stringPointer("bob@example.com")); err != nil {
		t.Fatalf("expected mail change success, got: %v", err)
	}

	if person.Name().GetValue() != "bob" {
		t.Fatalf("expected renamed person, got: %s", person.Name().GetValue())
	}
	if person.MailAddress().GetValue() != "bob@example.com" {
		t.Fatalf("expected changed mail address, got: %s", person.MailAddress().GetValue())
	}
}

func TestPersonEnsureReadyToUpdateRequiresLifecycleState(t *testing.T) {
	person, err := NewPerson(&NewPersonArgs{
		Name:        stringPointer("alice"),
		MailAddress: stringPointer("alice@example.com"),
	})
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	if err := person.EnsureReadyToUpdate(); err == nil {
		t.Fatal("expected update lifecycle error")
	}
}

func TestReconstructPersonListRequiresIdentityForEachPerson(t *testing.T) {
	_, err := ReconstructPersonList(&NewPersonListArgs{
		Content: []NewPersonArgs{
			{
				ID:          intPointer(1),
				Name:        stringPointer("alice"),
				MailAddress: stringPointer("alice@example.com"),
			},
			{
				Name:        stringPointer("bob"),
				MailAddress: stringPointer("bob@example.com"),
			},
		},
	})
	if err == nil {
		t.Fatal("expected identity error")
	}
}

func TestPersonListCanManageCollectionRules(t *testing.T) {
	person, err := ReconstructPerson(&NewPersonArgs{
		ID:          intPointer(1),
		Name:        stringPointer("alice"),
		MailAddress: stringPointer("alice@example.com"),
	})
	if err != nil {
		t.Fatalf("failed to reconstruct person: %v", err)
	}

	personList, err := ReconstructPersonList(&NewPersonListArgs{
		Content: []NewPersonArgs{
			{
				ID:          intPointer(2),
				Name:        stringPointer("bob"),
				MailAddress: stringPointer("bob@example.com"),
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to reconstruct person list: %v", err)
	}

	if personList.IsEmpty() {
		t.Fatal("expected person list not to be empty")
	}
	if personList.Count() != 1 {
		t.Fatalf("expected person count 1, got %d", personList.Count())
	}
	if err := personList.Append(*person); err != nil {
		t.Fatalf("expected append success, got: %v", err)
	}
	if !personList.ContainsIdentity(person.Identity()) {
		t.Fatal("expected appended identity")
	}
	if !personList.ContainsMailAddress(person.MailAddress()) {
		t.Fatal("expected appended mail address")
	}
	if err := personList.Append(*person); err == nil {
		t.Fatal("expected duplicate person error")
	}
}

func intPointer(value int) *int {
	return &value
}

func stringPointer(value string) *string {
	return &value
}

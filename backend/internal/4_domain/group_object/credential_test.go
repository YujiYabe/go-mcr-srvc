package group_object

import "testing"

func TestCredentialCanAuthenticate(
	t *testing.T,
) {
	credential, err := NewCredential(&NewCredentialArgs{
		ClientID:     stringPointer("client-id"),
		ClientSecret: stringPointer("client-secret"),
	})
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	if !credential.CanAuthenticate() {
		t.Fatal("expected credential to be ready for authentication")
	}
	if err := credential.EnsureReadyToAuthenticate(); err != nil {
		t.Fatalf("expected credential to authenticate, got: %v", err)
	}
}

func TestCredentialRequiresClientSecret(
	t *testing.T,
) {
	credential, err := NewCredential(&NewCredentialArgs{
		ClientID: stringPointer("client-id"),
	})
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	if credential.CanAuthenticate() {
		t.Fatal("expected credential not to authenticate without secret")
	}
	if err := credential.EnsureReadyToAuthenticate(); err == nil {
		t.Fatal("expected client secret error")
	}
}

func TestCredentialCanRotateSecret(
	t *testing.T,
) {
	credential, err := NewCredential(&NewCredentialArgs{
		ClientID:     stringPointer("client-id"),
		ClientSecret: stringPointer("old-secret"),
	})
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	if err := credential.RotateSecret(stringPointer("new-secret")); err != nil {
		t.Fatalf("expected secret rotation success, got: %v", err)
	}
	if credential.ClientSecret().GetValue() != "new-secret" {
		t.Fatalf("expected rotated secret, got: %s", credential.ClientSecret().GetValue())
	}
}

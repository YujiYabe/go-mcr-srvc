package usecase

import (
	"context"
	"errors"
	"strings"
	"testing"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type fakeGatewayDB struct {
	runInTransactionCalled bool
	updatePersonCalled     bool
	getByConditionCalled   bool
}

func (receiver *fakeGatewayDB) RunInTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	receiver.runInTransactionCalled = true
	return fn(ctx)
}

func (receiver *fakeGatewayDB) GetPersonList(
	ctx context.Context,
) (
	groupObject.PersonList,
	error,
) {
	return groupObject.PersonList{}, nil
}

func (receiver *fakeGatewayDB) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	groupObject.PersonList,
	error,
) {
	receiver.getByConditionCalled = true
	return groupObject.PersonList{}, nil
}

func (receiver *fakeGatewayDB) UpdatePerson(
	ctx context.Context,
	newPerson groupObject.Person,
) error {
	receiver.updatePersonCalled = true
	return nil
}

type fakeGatewayExternal struct {
	fetchAccessTokenCalled bool
}

func (receiver *fakeGatewayExternal) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	typeObject.AccessToken,
	error,
) {
	receiver.fetchAccessTokenCalled = true
	return typeObject.NewAccessToken(stringPointer("access-token"))
}

func (receiver *fakeGatewayExternal) ViaGRPC(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	groupObject.PersonList,
	error,
) {
	return groupObject.PersonList{}, nil
}

func (receiver *fakeGatewayExternal) PublishTestTopic(
	ctx context.Context,
) error {
	return nil
}

func TestGetPersonListByConditionRequiresCondition(t *testing.T) {
	dbGateway := &fakeGatewayDB{}
	useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})
	person := newTestPerson(t, nil, nil, nil)

	_, err := useCase.GetPersonListByCondition(context.Background(), person)
	if err == nil {
		t.Fatal("expected condition error")
	}
	if !strings.Contains(err.Error(), "person search condition is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if dbGateway.getByConditionCalled {
		t.Fatal("gateway should not be called when condition is empty")
	}
}

func TestFetchAccessTokenRequiresCredential(t *testing.T) {
	externalGateway := &fakeGatewayExternal{}
	useCase := NewUseCase(nil, &fakeGatewayDB{}, externalGateway)
	credential := newTestCredential(t, "", "secret")

	_, err := useCase.FetchAccessToken(context.Background(), credential)
	if err == nil {
		t.Fatal("expected credential error")
	}
	if !strings.Contains(err.Error(), "client id is required") {
		t.Fatalf("unexpected error: %v", err)
	}
	if externalGateway.fetchAccessTokenCalled {
		t.Fatal("gateway should not be called when credential is invalid")
	}
}

func TestUpdatePersonRunsInTransaction(t *testing.T) {
	dbGateway := &fakeGatewayDB{}
	useCase := NewUseCase(nil, dbGateway, &fakeGatewayExternal{})
	person := newTestPerson(t, intPointer(1), stringPointer("name"), stringPointer("test@example.com"))

	if err := useCase.UpdatePerson(context.Background(), person); err != nil {
		t.Fatalf("expected update success, got: %v", err)
	}
	if !dbGateway.runInTransactionCalled {
		t.Fatal("transaction boundary was not used")
	}
	if !dbGateway.updatePersonCalled {
		t.Fatal("update gateway was not called")
	}
}

func TestEnsureContextReadyReturnsCanceledError(t *testing.T) {
	useCase := NewUseCase(nil, &fakeGatewayDB{}, &fakeGatewayExternal{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := useCase.GetPersonList(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got: %v", err)
	}
}

func newTestPerson(
	t *testing.T,
	id *int,
	name *string,
	mailAddress *string,
) groupObject.Person {
	t.Helper()

	person, err := groupObject.NewPerson(&groupObject.NewPersonArgs{
		ID:          id,
		Name:        name,
		MailAddress: mailAddress,
	})
	if err != nil {
		t.Fatalf("failed to create person: %v", err)
	}

	return *person
}

func newTestCredential(
	t *testing.T,
	clientID string,
	clientSecret string,
) groupObject.Credential {
	t.Helper()

	credential, err := groupObject.NewCredential(&groupObject.NewCredentialArgs{
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
	})
	if err != nil {
		t.Fatalf("failed to create credential: %v", err)
	}

	return *credential
}

func intPointer(value int) *int {
	return &value
}

func stringPointer(value string) *string {
	return &value
}

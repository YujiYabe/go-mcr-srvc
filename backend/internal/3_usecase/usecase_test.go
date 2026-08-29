package usecase

import (
	"context"
	"errors"
	"testing"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type fakeGatewayDB struct {
	runInTransactionCalled     bool
	updateUserCalled           bool
	updateEmploymentCalled     bool
	getListCalled              bool
	getByConditionCalled       bool
	getValidationWordsCalled   bool
	addValidationWordCalled    bool
	updateValidationWordCalled bool
	deleteValidationWordCalled bool
	calls                      []string
	getListErr                 error
	getByConditionErr          error
	updateUserErr              error
	updateUserEmploymentErr    error
	getValidationWordsErr      error
	validationWordUpdateErr    error
	validationWords            []string
}

func (receiver *fakeGatewayDB) RunInTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	receiver.runInTransactionCalled = true
	return fn(ctx)
}

func (receiver *fakeGatewayDB) GetUserList(
	_ context.Context,
) (
	groupObject.UserList,
	error,
) {
	receiver.getListCalled = true
	return groupObject.UserList{}, receiver.getListErr
}

func (receiver *fakeGatewayDB) GetUserListByCondition(
	_ context.Context,
	_ groupObject.User,
) (
	groupObject.UserList,
	error,
) {
	receiver.getByConditionCalled = true
	return groupObject.UserList{}, receiver.getByConditionErr
}

func (receiver *fakeGatewayDB) UpdateUser(
	_ context.Context,
	_ groupObject.User,
) error {
	receiver.updateUserCalled = true
	receiver.calls = append(receiver.calls, "update_user")
	return receiver.updateUserErr
}

func (receiver *fakeGatewayDB) UpdateUserEmployment(
	_ context.Context,
	_ groupObject.UserEmployment,
) error {
	receiver.updateEmploymentCalled = true
	receiver.calls = append(receiver.calls, "update_user_employment")
	return receiver.updateUserEmploymentErr
}

func (receiver *fakeGatewayDB) GetValidationWords(
	_ context.Context,
	_ string,
	_ bool,
) (
	[]string,
	error,
) {
	receiver.getValidationWordsCalled = true
	return receiver.validationWords, receiver.getValidationWordsErr
}

func (receiver *fakeGatewayDB) AddValidationWord(
	_ context.Context,
	_ string,
	_ bool,
	_ string,
) error {
	receiver.addValidationWordCalled = true
	return receiver.validationWordUpdateErr
}

func (receiver *fakeGatewayDB) UpdateValidationWord(
	_ context.Context,
	_ string,
	_ bool,
	_ string,
	_ string,
) error {
	receiver.updateValidationWordCalled = true
	return receiver.validationWordUpdateErr
}

func (receiver *fakeGatewayDB) DeleteValidationWord(
	_ context.Context,
	_ string,
	_ bool,
	_ string,
) error {
	receiver.deleteValidationWordCalled = true
	return receiver.validationWordUpdateErr
}

type fakeGatewayExternal struct {
	fetchAccessTokenCalled bool
	viaGRPCCalled          bool
	publishTestTopicCalled bool
	fetchAccessTokenErr    error
	viaGRPCErr             error
	publishTestTopicErr    error
}

func (receiver *fakeGatewayExternal) FetchAccessToken(
	_ context.Context,
	_ groupObject.Credential,
) (
	typeObject.AccessToken,
	error,
) {
	receiver.fetchAccessTokenCalled = true
	if receiver.fetchAccessTokenErr != nil {
		return typeObject.AccessToken{}, receiver.fetchAccessTokenErr
	}
	return typeObject.NewAccessToken(stringPointer("access-token"))
}

func (receiver *fakeGatewayExternal) GetUserViaGRPC(
	_ context.Context,
	_ groupObject.User,
) (
	groupObject.UserList,
	error,
) {
	receiver.viaGRPCCalled = true
	return groupObject.UserList{}, receiver.viaGRPCErr
}

func (receiver *fakeGatewayExternal) PublishTestTopic(
	_ context.Context,
) error {
	receiver.publishTestTopicCalled = true
	return receiver.publishTestTopicErr
}

func TestEnsureContextReadyReturnsCanceledError(t *testing.T) {
	useCase := NewUseCase(nil, &fakeGatewayDB{}, &fakeGatewayExternal{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := useCase.GetUserList(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got: %v", err)
	}
}

func newTestUser(
	t *testing.T,
	id *int,
	name *string,
	email *string,
) groupObject.User {
	t.Helper()

	user, err := groupObject.NewUser(&groupObject.NewUserArgs{
		ID:    id,
		Name:  name,
		Email: email,
	})
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	return *user
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

func newTestUserEmployment(
	t *testing.T,
	userID int,
	isPrimary bool,
) groupObject.UserEmployment {
	t.Helper()

	employeeCode := "EMP001"
	employmentType := "full_time"
	userEmployment, err := groupObject.NewUserEmployment(&groupObject.NewUserEmploymentArgs{
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

	return *userEmployment
}

func intPointer(value int) *int {
	return &value
}

func stringPointer(value string) *string {
	return &value
}

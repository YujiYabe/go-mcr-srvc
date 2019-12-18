package mock_usecase

import (
	reflect "reflect"
	"testing"

	domain "app/domain"
	"app/usecase"

	"github.com/golang/mock/gomock"
)

func TestView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected domain.Users
	var err error

	mockSample := NewMockUCUserRepository(ctrl)
	mockSample.EXPECT().IFDBFindAll().Return(expected, err)

	// mockを利用してtodoUsecase.View()をテストする
	// todoUsecase := usecase.NewTodoUsecase(mockSample)
	userUsecase := &usecase.UCUserInteractor{mockSample}
	result, err := userUsecase.UCUIUsers()

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual FindAll() is not same as expected")
	}

}

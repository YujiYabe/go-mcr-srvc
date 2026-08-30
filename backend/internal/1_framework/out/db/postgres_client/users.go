package postgres_client

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"backend/internal/1_framework/out/db/postgres_client/models"
	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
	"backend/internal/logger"
)

func (receiver *PostgresClient) ReplaceUser(
	ctx context.Context,
	name string,
	email string,
	userID string,
) (
	err error,
) {
	err = receiver.conn(ctx).Transaction(func(tx *gorm.DB) error {
		err := receiver.AddUser(tx, name, email)
		if err != nil {
			return err
		}
		err = receiver.DeleteUser(tx, userID)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (receiver *PostgresClient) AddUser(
	tx *gorm.DB,
	name string,
	email string,
) (
	err error,
) {
	record := models.User{
		FullName: sql.NullString{String: name, Valid: name != ""},
		Email:    email,
	}

	return tx.
		Omit("Auth0UserID", "CreatedAt", "UpdatedAt").
		Create(&record).
		Error
}

func (receiver *PostgresClient) DeleteUser(
	tx *gorm.DB,
	userID string,
) (
	err error,
) {
	return tx.Delete(&models.User{}, userID).Error
}

func (receiver *PostgresClient) GetUserList(
	ctx context.Context,
) (
	userList groupObject.UserList,
	err error,
) {
	users := []models.User{} // SQL結果保存用

	result := receiver.conn(ctx).
		Model(&models.User{}).
		Select("id", "email", "full_name").
		Find(&users)

	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected == 0 {
		return
	}

	userArgs := make([]groupObject.NewUserArgs, 0, len(users))
	for _, user := range users {
		userArgs = append(userArgs, groupObject.NewUserArgs{
			ID:    &user.ID,
			Name:  stringFromNullString(user.FullName),
			Email: &user.Email,
		})
	}

	return groupObject.ReconstructUserList(&groupObject.NewUserListArgs{
		Content: userArgs,
	})
}

func (receiver *PostgresClient) GetUser(
	ctx context.Context,
	id typeObject.ID,
) (
	user groupObject.User,
	err error,
) {
	user = groupObject.User{}   // ドメインロジック用
	resultUser := models.User{} // SQL結果保存用

	result := receiver.conn(ctx).
		Model(&models.User{}).
		Select("id", "email", "full_name").
		Where(&models.User{ID: id.GetValue()}).
		Take(&resultUser)

	if result.Error != nil {
		err = result.Error
		return
	}

	args := &groupObject.NewUserArgs{
		ID:    &resultUser.ID,
		Name:  stringFromNullString(resultUser.FullName),
		Email: &resultUser.Email,
	}
	newUser, err := groupObject.ReconstructUser(args)
	if err != nil {
		return
	}

	return *newUser, nil
}

func (receiver *PostgresClient) UpdateUser(
	ctx context.Context,
	newUser groupObject.User,
) (
	err error,
) {
	if err := newUser.EnsureReadyToUpdate(); err != nil {
		return err
	}

	record := models.User{
		FullName: sql.NullString{
			String: newUser.Name().GetValue(),
			Valid:  true,
		},
		Email: newUser.Email().GetValue(),
	}

	result := receiver.conn(ctx).
		Model(&models.User{}).
		Where(&models.User{ID: newUser.Identity().GetValue()}).
		Select("full_name", "email").
		Updates(&record)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetUserListByCondition ...
func (receiver *PostgresClient) GetUserListByCondition(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	// logger.Logging(
	// 	ctx,
	// 	middlewareRequestContext.GetRequestContext(ctx).TraceID.GetValue(),
	// )

	users := []models.User{} // SQL結果保存用

	conn := receiver.conn(ctx).
		Model(&models.User{}).
		Select("id", "email", "full_name")

	if !reqUser.Email().GetIsNil() && reqUser.Email().GetValue() != "" {
		conn = conn.Where("email = ?", reqUser.Email().GetValue())
	}

	if !reqUser.Name().GetIsNil() && reqUser.Name().GetValue() != "" {
		conn = conn.Where("full_name LIKE ?", "%"+reqUser.Name().GetValue()+"%")
	}

	result := conn.Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}

	userArgs := make([]groupObject.NewUserArgs, 0, len(users))
	for _, user := range users {
		userArgs = append(userArgs, groupObject.NewUserArgs{
			ID:    &user.ID,
			Name:  stringFromNullString(user.FullName),
			Email: &user.Email,
		})
	}

	// logger.Logging(
	// 	ctx,
	// 	middlewareRequestContext.GetRequestContext(ctx).TraceID.GetValue(),
	// )

	resUserList, err = groupObject.ReconstructUserList(&groupObject.NewUserListArgs{
		Content: userArgs,
	})
	if err != nil {
		logger.Logging(ctx, err)
		return
	}

	return
}

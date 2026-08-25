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

type txContextKey struct{}

type userRecord struct {
	ID       int
	Email    string
	FullName sql.NullString
}

func (receiver *PostgresClient) RunInTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	return receiver.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(context.WithValue(ctx, txContextKey{}, tx))
	})
}

func (receiver *PostgresClient) conn(
	ctx context.Context,
) *gorm.DB {
	if tx, ok := ctx.Value(txContextKey{}).(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return receiver.Conn.WithContext(ctx)
}

func (receiver *PostgresClient) ReplaceUser(
	ctx context.Context,
	name string,
	email string,
	id string,
) error {
	err := receiver.conn(ctx).Transaction(func(tx *gorm.DB) error {
		err := receiver.AddUser(tx, name, email)
		if err != nil {
			return err
		}
		err = receiver.DeleteUser(tx, id)
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
) error {
	return tx.Table("users").Create(map[string]interface{}{
		"full_name": sql.NullString{String: name, Valid: name != ""},
		"email":     email,
	}).Error
}

func (receiver *PostgresClient) DeleteUser(
	tx *gorm.DB,
	id string,
) error {
	return tx.Delete(&models.User{}, id).Error
}

func (receiver *PostgresClient) GetUserList(
	ctx context.Context,
) (
	userList groupObject.UserList,
	err error,
) {
	users := []userRecord{} // SQL結果保存用

	result := receiver.conn(ctx).
		Table("users").
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
	user = groupObject.User{}  // ドメインロジック用
	resultUser := userRecord{} // SQL結果保存用

	result := receiver.conn(ctx).
		Table("users").
		Select("id", "email", "full_name").
		Where("id = ?", id.GetValue()).
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
) error {
	if err := newUser.EnsureReadyToUpdate(); err != nil {
		return err
	}

	result := receiver.conn(ctx).
		Table("users").
		Where("id = ?", newUser.Identity().GetValue()).
		Updates(map[string]interface{}{
			"full_name": newUser.Name().GetValue(),
			"email":     newUser.Email().GetValue(),
		})
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
	// 	requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue(),
	// )

	users := []userRecord{} // SQL結果保存用

	conn := receiver.conn(ctx).
		Table("users").
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
	// 	requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue(),
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

func stringFromNullString(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}

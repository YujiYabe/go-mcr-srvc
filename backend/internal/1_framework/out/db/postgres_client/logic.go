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

func (receiver *PostgresClient) ReplacePerson(
	ctx context.Context,
	name string,
	email string,
	id string,
) error {
	err := receiver.conn(ctx).Transaction(func(tx *gorm.DB) error {
		err := receiver.AddPerson(tx, name, email)
		if err != nil {
			return err
		}
		err = receiver.DeletePerson(tx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (receiver *PostgresClient) AddPerson(
	tx *gorm.DB,
	name string,
	email string,
) error {
	return tx.Table("users").Create(map[string]interface{}{
		"full_name": sql.NullString{String: name, Valid: name != ""},
		"email":     email,
	}).Error
}

func (receiver *PostgresClient) DeletePerson(
	tx *gorm.DB,
	id string,
) error {
	return tx.Delete(&models.User{}, id).Error
}

func (receiver *PostgresClient) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
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

	personArgs := make([]groupObject.NewPersonArgs, 0, len(users))
	for _, user := range users {
		personArgs = append(personArgs, groupObject.NewPersonArgs{
			ID:          &user.ID,
			Name:        stringFromNullString(user.FullName),
			MailAddress: &user.Email,
		})
	}

	return groupObject.ReconstructPersonList(&groupObject.NewPersonListArgs{
		Content: personArgs,
	})
}

func (receiver *PostgresClient) GetPerson(
	ctx context.Context,
	id typeObject.ID,
) (
	person groupObject.Person,
	err error,
) {
	person = groupObject.Person{} // ドメインロジック用
	resultUser := userRecord{}    // SQL結果保存用

	result := receiver.conn(ctx).
		Table("users").
		Select("id", "email", "full_name").
		Where("id = ?", id.GetValue()).
		Take(&resultUser)

	if result.Error != nil {
		err = result.Error
		return
	}

	args := &groupObject.NewPersonArgs{
		ID:          &resultUser.ID,
		Name:        stringFromNullString(resultUser.FullName),
		MailAddress: &resultUser.Email,
	}
	newPerson, err := groupObject.ReconstructPerson(args)
	if err != nil {
		return
	}

	return *newPerson, nil
}

func (receiver *PostgresClient) UpdatePerson(
	ctx context.Context,
	newPerson groupObject.Person,
) error {
	if err := newPerson.EnsureReadyToUpdate(); err != nil {
		return err
	}

	result := receiver.conn(ctx).
		Table("users").
		Where("id = ?", newPerson.Identity().GetValue()).
		Updates(map[string]interface{}{
			"full_name": newPerson.Name().GetValue(),
			"email":     newPerson.MailAddress().GetValue(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetPersonListByCondition ...
func (receiver *PostgresClient) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
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

	if !reqPerson.MailAddress().GetIsNil() && reqPerson.MailAddress().GetValue() != "" {
		conn = conn.Where("email = ?", reqPerson.MailAddress().GetValue())
	}

	if !reqPerson.Name().GetIsNil() && reqPerson.Name().GetValue() != "" {
		conn = conn.Where("full_name LIKE ?", "%"+reqPerson.Name().GetValue()+"%")
	}

	result := conn.Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}

	personArgs := make([]groupObject.NewPersonArgs, 0, len(users))
	for _, user := range users {
		personArgs = append(personArgs, groupObject.NewPersonArgs{
			ID:          &user.ID,
			Name:        stringFromNullString(user.FullName),
			MailAddress: &user.Email,
		})
	}

	// logger.Logging(
	// 	ctx,
	// 	requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue(),
	// )

	resPersonList, err = groupObject.ReconstructPersonList(&groupObject.NewPersonListArgs{
		Content: personArgs,
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

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
	newUser := models.Person{
		Name:        sql.NullString{String: name, Valid: true},
		MailAddress: sql.NullString{String: email, Valid: true},
	}
	return tx.Create(&newUser).Error
}

func (receiver *PostgresClient) DeletePerson(
	tx *gorm.DB,
	id string,
) error {
	return tx.Delete(&models.Person{}, id).Error
}

func (receiver *PostgresClient) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
	err error,
) {
	persons := []models.Person{} // SQL結果保存用

	result := receiver.conn(ctx).
		Table("persons").
		Find(&persons)

	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected == 0 {
		return
	}

	personArgs := make([]groupObject.NewPersonArgs, 0, len(persons))
	for _, person := range persons {
		personArgs = append(personArgs, groupObject.NewPersonArgs{
			ID:          &person.ID,
			Name:        &person.Name.String,
			MailAddress: &person.MailAddress.String,
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
	person = groupObject.Person{}   // ドメインロジック用
	resultPerson := models.Person{} // SQL結果保存用

	result := receiver.conn(ctx).
		Table("persons").
		Where("id = ?", id.GetValue()).
		Take(&resultPerson)

	if result.Error != nil {
		err = result.Error
		return
	}

	args := &groupObject.NewPersonArgs{
		ID:          &resultPerson.ID,
		Name:        &resultPerson.Name.String,
		MailAddress: &resultPerson.MailAddress.String,
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
		Table("persons").
		Where("id = ?", newPerson.Identity().GetValue()).
		Updates(map[string]interface{}{
			"name":         newPerson.Name().GetValue(),
			"mail_address": newPerson.MailAddress().GetValue(),
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

	persons := []models.Person{} // SQL結果保存用

	conn := receiver.conn(ctx).Table("persons")

	if !reqPerson.MailAddress().GetIsNil() && reqPerson.MailAddress().GetValue() != "" {
		conn = conn.Where("mail_address = ?", reqPerson.MailAddress().GetValue())
	}

	if !reqPerson.Name().GetIsNil() && reqPerson.Name().GetValue() != "" {
		conn = conn.Where("name LIKE ?", "%"+reqPerson.Name().GetValue()+"%")
	}

	result := conn.Find(&persons)
	if result.Error != nil {
		err = result.Error
		return
	}

	personArgs := make([]groupObject.NewPersonArgs, 0, len(persons))
	for _, person := range persons {
		personArgs = append(personArgs, groupObject.NewPersonArgs{
			ID:          &person.ID,
			Name:        &person.Name.String,
			MailAddress: &person.MailAddress.String,
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

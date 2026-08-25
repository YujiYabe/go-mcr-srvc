package postgres_client

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	"backend/internal/1_framework/out/db/postgres_client/models"
	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
	"backend/internal/logger"
)

func (receiver *PostgresClient) WithOutTx(
	ctx context.Context,
) (
	tx *gorm.DB,
) {
	return receiver.Conn.WithContext(ctx)
}

func (receiver *PostgresClient) BeginTx(
	ctx context.Context,
) (
	tx *gorm.DB,
) {
	return receiver.Conn.WithContext(ctx).Begin()
}

func (receiver *PostgresClient) EndTx(
	ctx context.Context,
	tx *gorm.DB,
	isSuccess bool,
) (
	err error,
) {
	if isSuccess {
		if err = tx.Commit().Error; err != nil {
			logger.Logging(ctx, fmt.Errorf("commit failed: %w", err))
		}
	} else {
		if err = tx.Rollback().Error; err != nil {
			logger.Logging(ctx, fmt.Errorf("rollback failed: %w", err))
		}
	}
	return
}

func (receiver *PostgresClient) ReplacePerson(
	ctx context.Context,
	name string,
	email string,
	id string,
) error {
	err := receiver.Conn.Transaction(func(tx *gorm.DB) error {
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
	personList = groupObject.PersonList{} // ドメインロジック用
	persons := []models.Person{}          // SQL結果保存用

	result := receiver.Conn.WithContext(ctx).
		Table("persons").
		Find(&persons)

	if result.Error != nil {
		err = result.Error
		return
	}

	if result.RowsAffected == 0 {
		return
	}

	for _, person := range persons {
		args := &groupObject.NewPersonArgs{
			ID:          &person.ID,
			Name:        &person.Name.String,
			MailAddress: &person.MailAddress.String,
		}
		domainPerson, createErr := groupObject.NewPerson(args)
		if createErr != nil {
			err = createErr
			return
		}

		personList.Append(*domainPerson)
	}

	return
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

	result := receiver.Conn.WithContext(ctx).
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
	newPerson, err := groupObject.NewPerson(args)
	if err != nil {
		return
	}

	return *newPerson, nil
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

	resPersonList = groupObject.PersonList{} // ドメインロジック用
	persons := []models.Person{}             // SQL結果保存用

	conn := receiver.Conn.WithContext(ctx).Table("persons")

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

	for _, person := range persons {
		args := &groupObject.NewPersonArgs{
			ID:          &person.ID,
			Name:        &person.Name.String,
			MailAddress: &person.MailAddress.String,
		}
		domainPerson, createErr := groupObject.NewPerson(args)
		if createErr != nil {
			err = createErr
			logger.Logging(ctx, err)
			return
		}

		resPersonList.Append(*domainPerson)
	}

	// logger.Logging(
	// 	ctx,
	// 	requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue(),
	// )

	return
}

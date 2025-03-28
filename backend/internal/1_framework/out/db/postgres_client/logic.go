package postgres_client

import (
	"context"
	// sql.NullStringを使用するために追加
	// gormを使用するために追加

	"backend/internal/1_framework/out/db/postgres_client/models"
	groupObject "backend/internal/4_domain/group_object"
	"backend/internal/logger"

	"database/sql" // sql.NullStringを使用するために追加

	"gorm.io/gorm" // gormを使用するために追加
)

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
	if err := tx.Create(&newUser).Error; err != nil {
		return err // エラーが発生した場合はロールバック
	}
	return nil // 正常終了の場合はコミット
}

func (receiver *PostgresClient) DeletePerson(
	tx *gorm.DB,
	id string,
) error {
	if err := tx.Delete(&models.Person{}, id).Error; err != nil {
		return err // エラーが発生した場合はロールバック
	}
	return nil // 正常終了の場合はコミット
}

func (receiver *PostgresClient) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
) {
	personList = groupObject.PersonList{} // ドメインロジック用
	persons := []models.Person{}          // SQL結果保存用

	result := receiver.Conn.
		Table("persons").
		Find(&persons)

	if result.Error != nil {
		personList.SetError(ctx, result.Error)
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
		person := groupObject.NewPerson(ctx, args)

		if person.GetError() != nil {
			personList.SetError(ctx, person.GetError())
			return
		}

		personList.Content = append(
			personList.Content,
			*person,
		)
	}

	return
}

// GetPersonListByCondition ...
func (receiver *PostgresClient) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	// logger.Logging(
	// 	ctx,
	// 	groupObject.GetRequestContext(ctx).TraceID.GetValue(),
	// )

	resPersonList = groupObject.PersonList{} // ドメインロジック用
	persons := []models.Person{}             // SQL結果保存用

	conn := receiver.Conn.Table("persons")

	if !reqPerson.MailAddress.GetIsNil() && reqPerson.MailAddress.GetValue() != "" {
		conn.Where("mail_address = ?", reqPerson.MailAddress.GetValue())
	}

	if !reqPerson.Name.GetIsNil() && reqPerson.Name.GetValue() != "" {
		conn.Where("name LIKE ?", "%"+reqPerson.Name.GetValue()+"%")
	}

	result := conn.Find(&persons)
	if result.Error != nil {
		resPersonList.SetError(ctx, result.Error)
		return
	}

	for _, person := range persons {
		args := &groupObject.NewPersonArgs{
			ID:          &person.ID,
			Name:        &person.Name.String,
			MailAddress: &person.MailAddress.String,
		}
		person := groupObject.NewPerson(ctx, args)

		if person.GetError() != nil {
			logger.Logging(ctx, person.GetError())
			resPersonList.SetError(ctx, person.GetError())
			return
		}

		resPersonList.Content = append(
			resPersonList.Content,
			*person,
		)
	}

	// logger.Logging(
	// 	ctx,
	// 	groupObject.GetRequestContext(ctx).TraceID.GetValue(),
	// )

	return
}

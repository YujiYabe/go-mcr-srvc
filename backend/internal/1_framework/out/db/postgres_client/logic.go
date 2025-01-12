package postgres_client

import (
	"context"

	"backend/internal/1_framework/out/db/postgres_client/models"
	groupObject "backend/internal/4_domain/group_object"
	"backend/pkg"
)

// GetPersonList ...
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
	pkg.Logging(
		ctx,
		groupObject.GetRequestContext(ctx).TraceID.GetValue(),
	)

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
			pkg.Logging(ctx, person.GetError())
			resPersonList.SetError(ctx, person.GetError())
			return
		}

		resPersonList.Content = append(
			resPersonList.Content,
			*person,
		)
	}

	pkg.Logging(
		ctx,
		groupObject.GetRequestContext(ctx).TraceID.GetValue(),
	)

	return
}

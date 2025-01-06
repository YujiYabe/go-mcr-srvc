package postgres_client

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/1_framework/out/db/postgres_client/models"
	"backend/internal/2_adapter/gateway"
	groupObject "backend/internal/4_domain/group_object"
	"backend/pkg"
)

type (
	// PostgresClient ...
	PostgresClient struct {
		Conn *gorm.DB
	}
)

// NewToPostgres ...
func NewToPostgres() gateway.ToPostgres {
	ctx := context.Background()
	conn, err := open(30)
	if err != nil {
		pkg.Logging(ctx, err)
		panic(err)
	}

	postgresClient := new(PostgresClient)
	postgresClient.Conn = conn
	return postgresClient
}

func open(count uint) (*gorm.DB, error) {
	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(pkg.PostgresDSN), &gorm.Config{})

	if err != nil {
		if count == 0 {
			pkg.Logging(ctx, err)
			return nil, fmt.Errorf(
				"retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(count)
	}

	return db, nil
}

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
		pkg.Logging(ctx, result.Error)
		personList.SetError(ctx, result.Error)
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

// GetPersonByCondition ...
func (receiver *PostgresClient) GetPersonByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
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

	return
}

package postgres

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/1_framework/db/postgres/models"
	"backend/internal/2_adapter/gateway"
	"backend/internal/4_domain/struct_object"
	"backend/pkg"
)

type (
	// Postgres ...
	Postgres struct {
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

	postgres := new(Postgres)
	postgres.Conn = conn
	return postgres
}

func open(count uint) (*gorm.DB, error) {
	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(pkg.PostgresDSN), &gorm.Config{})

	if err != nil {
		if count == 0 {
			pkg.Logging(ctx, err)
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(count)
	}

	return db, nil
}

// GetPersonList ...
func (receiver *Postgres) GetPersonList(
	ctx context.Context,
) (
	personList struct_object.PersonList,
	err error,
) {
	personList = struct_object.PersonList{} // ドメインロジック用
	persons := []models.Person{}            // SQL結果保存用

	result := receiver.Conn.
		Table("persons").
		Find(&persons)

	if result.Error != nil {
		pkg.Logging(ctx, result.Error)
		return personList, result.Error
	}
	for _, person := range persons {
		args := &struct_object.NewPersonArgs{
			ID:          &person.ID,
			Name:        &person.Name.String,
			MailAddress: &person.MailAddress.String,
		}
		person := struct_object.NewPerson(args)

		if person.Err != nil {
			pkg.Logging(ctx, person.Err)
			return personList, person.Err
		}

		personList = append(personList, *person)
	}

	return
}

// GetPersonByCondition ...
func (receiver *Postgres) GetPersonByCondition(
	ctx context.Context,
	reqPerson struct_object.Person,
) (
	resPersonList struct_object.PersonList,
	err error,
) {
	resPersonList = struct_object.PersonList{} // ドメインロジック用
	persons := []models.Person{}               // SQL結果保存用

	conn := receiver.Conn.Table("persons")

	if !reqPerson.MailAddress.Content.IsNil && reqPerson.MailAddress.Content.GetValue() != "" {
		conn.Where("mail_address = ?", reqPerson.MailAddress.Content.GetValue())
	}

	if !reqPerson.Name.Content.IsNil && reqPerson.Name.Content.GetValue() != "" {
		conn.Where("name LIKE ?", "%"+reqPerson.Name.Content.GetValue()+"%")
	}

	result := conn.Find(&persons)
	if result.Error != nil {
		pkg.Logging(ctx, result.Error)
		return resPersonList, result.Error
	}
	for _, person := range persons {
		args := &struct_object.NewPersonArgs{
			ID:          &person.ID,
			Name:        &person.Name.String,
			MailAddress: &person.MailAddress.String,
		}
		person := struct_object.NewPerson(args)

		if person.Err != nil {
			pkg.Logging(ctx, person.Err)
			return resPersonList, person.Err
		}

		resPersonList = append(resPersonList, *person)
	}

	return
}

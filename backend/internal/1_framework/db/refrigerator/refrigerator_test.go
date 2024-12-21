package refrigerator

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUpdateVegetables(t *testing.T) {
	t.Run(
		"",
		func(t *testing.T) {
			// DBのモックを作成
			sqlDB, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer sqlDB.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: sqlDB,
			}), &gorm.Config{})

			refrigerator := Refrigerator{
				Conn: gormDB,
			}
			// 想定されるクエリとその引数に対して返す値の設定
			mock.ExpectBegin()
			mock.ExpectExec("UPDATE vegetables(.+)").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			mock.ExpectClose()

			ctx := context.Background()
			items := map[string]int{
				"aa": 1,
			}
			err = refrigerator.UpdateVegetables(ctx, items)

			if err != nil {
				t.Error(err.Error())
			}
		},
	)

}

package postgres_client

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetValidationWordsQueriesEnabledContainsRules(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql mock: %v", err)
	}
	defer sqlDB.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm mock db: %v", err)
	}

	rows := sqlmock.NewRows([]string{"word"}).
		AddRow("admin").
		AddRow("root")
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "word" FROM "validation_word_rules" WHERE ("validation_word_rules"."match_type" = $1 AND "validation_word_rules"."enabled" = $2) AND "validation_word_rules"."target_type" = $3 AND is_blacklist = $4 ORDER BY word ASC`,
	)).
		WithArgs("contains", true, "name", true).
		WillReturnRows(rows)

	client := &PostgresClient{Conn: gormDB}
	words, err := client.GetValidationWords(context.Background(), "name", true)
	if err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
	if !reflect.DeepEqual(words, []string{"admin", "root"}) {
		t.Fatalf("unexpected words: %v", words)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestAddValidationWordUpsertsContainsRule(t *testing.T) {
	client, mock, closeDB := newValidationWordRuleTestClient(t)
	defer closeDB()

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "validation_word_rules".*ON CONFLICT .* DO UPDATE.*RETURNING "id"`).
		WithArgs("name", true, "root", "contains", true, true, "contains").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	if err := client.AddValidationWord(context.Background(), "name", true, "root"); err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestUpdateValidationWordUpdatesContainsRule(t *testing.T) {
	client, mock, closeDB := newValidationWordRuleTestClient(t)
	defer closeDB()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "validation_word_rules"`)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := client.UpdateValidationWord(context.Background(), "name", true, "root", "admin"); err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestUpdateValidationWordReturnsNotFound(t *testing.T) {
	client, mock, closeDB := newValidationWordRuleTestClient(t)
	defer closeDB()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "validation_word_rules"`)).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	if err := client.UpdateValidationWord(context.Background(), "name", true, "missing", "admin"); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected record not found, got: %v", err)
	}
}

func TestDeleteValidationWordDeletesRule(t *testing.T) {
	client, mock, closeDB := newValidationWordRuleTestClient(t)
	defer closeDB()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "validation_word_rules"`)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := client.DeleteValidationWord(context.Background(), "name", true, "root"); err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestDeleteValidationWordReturnsNotFound(t *testing.T) {
	client, mock, closeDB := newValidationWordRuleTestClient(t)
	defer closeDB()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "validation_word_rules"`)).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	if err := client.DeleteValidationWord(context.Background(), "name", true, "missing"); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected record not found, got: %v", err)
	}
}

func newValidationWordRuleTestClient(
	t *testing.T,
) (
	*PostgresClient,
	sqlmock.Sqlmock,
	func(),
) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql mock: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm mock db: %v", err)
	}

	return &PostgresClient{Conn: gormDB}, mock, func() {
		_ = sqlDB.Close()
	}
}

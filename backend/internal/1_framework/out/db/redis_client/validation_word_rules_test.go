package redis_client

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redismock/v9"
)

func TestGetValidationWordsReturnsCachedWords(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{Conn: db}
	mock.ExpectGet("validation:word_rules:name:blacklist").SetVal(`["admin","root"]`)

	words, hit, err := client.GetValidationWords(context.Background(), "name", true)
	if err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
	if !hit {
		t.Fatal("expected cache hit")
	}
	if !reflect.DeepEqual(words, []string{"admin", "root"}) {
		t.Fatalf("unexpected words: %v", words)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet redis expectations: %v", err)
	}
}

func TestGetValidationWordsReturnsMiss(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{Conn: db}
	mock.ExpectGet("validation:word_rules:name:blacklist").RedisNil()

	words, hit, err := client.GetValidationWords(context.Background(), "name", true)
	if err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
	if hit {
		t.Fatal("expected cache miss")
	}
	if words != nil {
		t.Fatalf("expected no words, got: %v", words)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet redis expectations: %v", err)
	}
}

func TestGetValidationWordsReturnsDecodeError(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{Conn: db}
	mock.ExpectGet("validation:word_rules:name:blacklist").SetVal(`not-json`)

	_, hit, err := client.GetValidationWords(context.Background(), "name", true)
	if err == nil {
		t.Fatal("expected decode error")
	}
	if hit {
		t.Fatal("decode error should not be treated as cache hit")
	}
}

func TestSetAndDeleteValidationWords(t *testing.T) {
	db, mock := redismock.NewClientMock()
	client := &RedisClient{Conn: db}
	key := "validation:word_rules:name:blacklist"
	mock.ExpectSet(key, []byte(`["admin","root"]`), validationWordsCacheTTL).SetVal("OK")
	mock.ExpectDel(key).SetVal(1)

	if err := client.SetValidationWords(context.Background(), "name", true, []string{"admin", "root"}); err != nil {
		t.Fatalf("expected set success, got: %v", err)
	}
	if err := client.DeleteValidationWordsCache(context.Background(), "name", true); err != nil {
		t.Fatalf("expected delete success, got: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet redis expectations: %v", err)
	}
}

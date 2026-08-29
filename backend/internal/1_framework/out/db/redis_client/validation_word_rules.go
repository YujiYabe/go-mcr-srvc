package redis_client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	domain "backend/internal/4_domain"
)

func (receiver *RedisClient) GetValidationWords(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
) (
	words []string,
	hit bool,
	err error,
) {
	result, err := receiver.Conn.Get(ctx, validationWordsCacheKey(targetType, isBlacklist)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	if err := json.Unmarshal([]byte(result), &words); err != nil {
		return nil, false, err
	}

	return words, true, nil
}

func (receiver *RedisClient) SetValidationWords(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	words []string,
) error {
	value, err := json.Marshal(words)
	if err != nil {
		return err
	}

	return receiver.Conn.Set(
		ctx,
		validationWordsCacheKey(targetType, isBlacklist),
		value,
		0,
	).Err()
}

func (receiver *RedisClient) DeleteValidationWordsCache(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
) error {
	return receiver.Conn.Del(ctx, validationWordsCacheKey(targetType, isBlacklist)).Err()
}

func validationWordsCacheKey(
	targetType string,
	isBlacklist bool,
) string {
	return fmt.Sprintf(
		"validation:word_rules:%s:%s",
		targetType,
		domain.ValidationWordRuleTypeFromBlacklistFlag(isBlacklist),
	)
}

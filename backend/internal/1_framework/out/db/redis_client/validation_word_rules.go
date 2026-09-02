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
	words = nil
	result, err := receiver.Conn.Get(ctx, validationWordsCacheKey(targetType, isBlacklist)).Result()
	if errors.Is(err, redis.Nil) {
		words, hit, err = nil, false, nil
		return
	}
	if err != nil {
		words, hit = nil, false
		return
	}

	if err := json.Unmarshal([]byte(result), &words); err != nil {
		return nil, false, err
	}

	hit, err = true, nil
	return
}

func (receiver *RedisClient) SetValidationWords(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	words []string,
) (
	err error,
) {
	value, err := json.Marshal(words)
	if err != nil {
		return
	}

	err = receiver.Conn.Set(
		ctx,
		validationWordsCacheKey(targetType, isBlacklist),
		value,
		0,
	).Err()
	return
}

func (receiver *RedisClient) DeleteValidationWordsCache(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
) (
	err error,
) {
	err = receiver.Conn.Del(ctx, validationWordsCacheKey(targetType, isBlacklist)).Err()
	return
}

func validationWordsCacheKey(
	targetType string,
	isBlacklist bool,
) (
	key string,
) {
	key = fmt.Sprintf(
		"validation:word_rules:%s:%s",
		targetType,
		domain.ValidationWordRuleTypeFromBlacklistFlag(isBlacklist),
	)
	return
}

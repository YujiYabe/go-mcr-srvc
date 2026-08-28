package postgres_client

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type validationWordRuleRecord struct {
	Word string
}

func (receiver *PostgresClient) GetValidationWords(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
) (
	words []string,
	err error,
) {
	records := []validationWordRuleRecord{}

	result := receiver.conn(ctx).
		Table("validation_word_rules").
		Select("word").
		Where("target_type = ?", targetType).
		Where("is_blacklist = ?", isBlacklist).
		Where("enabled = ?", true).
		Where("match_type = ?", "contains").
		Order("word ASC").
		Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}

	words = make([]string, 0, len(records))
	for _, record := range records {
		words = append(words, record.Word)
	}

	return words, nil
}

func (receiver *PostgresClient) AddValidationWord(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	word string,
) error {
	return receiver.conn(ctx).
		Table("validation_word_rules").
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "target_type"},
				{Name: "is_blacklist"},
				{Name: "word"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"match_type": "contains",
				"enabled":    true,
				"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
			}),
		}).
		Create(map[string]interface{}{
			"target_type":  targetType,
			"is_blacklist": isBlacklist,
			"word":         word,
			"match_type":   "contains",
			"enabled":      true,
		}).Error
}

func (receiver *PostgresClient) UpdateValidationWord(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	oldWord string,
	newWord string,
) error {
	result := receiver.conn(ctx).
		Table("validation_word_rules").
		Where("target_type = ?", targetType).
		Where("is_blacklist = ?", isBlacklist).
		Where("word = ?", oldWord).
		Updates(map[string]interface{}{
			"word":       newWord,
			"match_type": "contains",
			"enabled":    true,
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (receiver *PostgresClient) DeleteValidationWord(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	word string,
) error {
	result := receiver.conn(ctx).
		Table("validation_word_rules").
		Where("target_type = ?", targetType).
		Where("is_blacklist = ?", isBlacklist).
		Where("word = ?", word).
		Delete(map[string]interface{}{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

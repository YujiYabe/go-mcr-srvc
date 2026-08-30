package postgres_client

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"backend/internal/1_framework/out/db/postgres_client/models"
	domain "backend/internal/4_domain"
)

func (receiver *PostgresClient) GetValidationWords(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
) (
	words []string,
	err error,
) {
	records := []models.ValidationWordRule{}

	result := receiver.conn(ctx).
		Model(&models.ValidationWordRule{}).
		Select("word").
		Scopes(validationWordRuleScope(targetType, isBlacklist)).
		Where(&models.ValidationWordRule{
			Enabled:   true,
			MatchType: domain.ValidationWordRuleMatchTypeContains,
		}).
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
) (
	err error,
) {
	record := models.ValidationWordRule{
		TargetType:  targetType,
		IsBlacklist: isBlacklist,
		Word:        word,
		MatchType:   domain.ValidationWordRuleMatchTypeContains,
		Enabled:     true,
	}

	return receiver.conn(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "target_type"},
				{Name: "is_blacklist"},
				{Name: "word"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"match_type": domain.ValidationWordRuleMatchTypeContains,
				"enabled":    true,
				"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
			}),
		}).
		Omit("CreatedAt", "UpdatedAt").
		Create(&record).
		Error
}

func (receiver *PostgresClient) UpdateValidationWord(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	oldWord string,
	newWord string,
) (
	err error,
) {
	result := receiver.conn(ctx).
		Model(&models.ValidationWordRule{}).
		Scopes(validationWordRuleScope(targetType, isBlacklist)).
		Where(&models.ValidationWordRule{Word: oldWord}).
		Updates(map[string]interface{}{
			"word":       newWord,
			"match_type": domain.ValidationWordRuleMatchTypeContains,
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
) (
	err error,
) {
	result := receiver.conn(ctx).
		Scopes(validationWordRuleScope(targetType, isBlacklist)).
		Where(&models.ValidationWordRule{Word: word}).
		Delete(&models.ValidationWordRule{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func validationWordRuleScope(
	targetType string,
	isBlacklist bool,
) (
	fn func(*gorm.DB) *gorm.DB,
) {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where(&models.ValidationWordRule{TargetType: targetType}).
			Where("is_blacklist = ?", isBlacklist)
	}
}

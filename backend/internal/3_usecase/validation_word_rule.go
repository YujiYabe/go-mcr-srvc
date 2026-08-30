package usecase

import (
	"context"
	"fmt"
)

func (receiver *useCase) GetValidationWords(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
) (
	words []string,
	err error,
) {
	if err := ensureContextReady(ctx, "GetValidationWords"); err != nil {
		return nil, err
	}

	words, err = receiver.ToGatewayDB.GetValidationWords(ctx, targetType, isBlacklist)
	if err != nil {
		return nil, fmt.Errorf("GetValidationWords: %w", err)
	}

	return words, nil
}

func (receiver *useCase) AddValidationWord(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	word string,
) (
	err error,
) {
	if err := ensureContextReady(ctx, "AddValidationWord"); err != nil {
		return err
	}

	if err := receiver.ToGatewayDB.AddValidationWord(
		ctx,
		targetType,
		isBlacklist,
		word,
	); err != nil {
		return fmt.Errorf("AddValidationWord: %w", err)
	}

	return nil
}

func (receiver *useCase) UpdateValidationWord(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	oldWord string,
	newWord string,
) (
	err error,
) {
	if err := ensureContextReady(ctx, "UpdateValidationWord"); err != nil {
		return err
	}

	if err := receiver.ToGatewayDB.UpdateValidationWord(
		ctx,
		targetType,
		isBlacklist,
		oldWord,
		newWord,
	); err != nil {
		return fmt.Errorf("UpdateValidationWord: %w", err)
	}

	return nil
}

func (receiver *useCase) DeleteValidationWord(
	ctx context.Context,
	targetType string,
	isBlacklist bool,
	word string,
) (
	err error,
) {
	if err := ensureContextReady(ctx, "DeleteValidationWord"); err != nil {
		return err
	}

	if err := receiver.ToGatewayDB.DeleteValidationWord(
		ctx,
		targetType,
		isBlacklist,
		word,
	); err != nil {
		return fmt.Errorf("DeleteValidationWord: %w", err)
	}

	return nil
}

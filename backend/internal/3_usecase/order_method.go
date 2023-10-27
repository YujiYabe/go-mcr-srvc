package usecase

import (
	"context"
	"mime/multipart"

	domain "backend/internal/4_domain"
)

// GetOrderList ...
func (receiver *useCase) GetOrderList(
	ctx context.Context,
) domain.OrderList {
	return receiver.ToDomain.GetOrderList(
		ctx,
	)
}

// GetReservingList ...
func (receiver *useCase) GetReservingList(
	ctx context.Context,
) domain.ReservingList {
	return receiver.ToDomain.GetReservingList(
		ctx,
	)
}

// GetReserving ...
func (receiver *useCase) GetReserving(
	ctx context.Context,
	number int,
) domain.Reserving {
	return receiver.ToDomain.GetReserving(
		ctx,
		number,
	)
}

// GetSoldList ...
func (receiver *useCase) GetSoldList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetSoldList(
		ctx,
	)
}

// SaveSold ...
func (receiver *useCase) SaveSold(
	ctx context.Context,
	newSold domain.Sold,
) {
	// 無効なJANcodeを排除
	janCodeList, languageCode := receiver.ToDomain.VerifyJANCodes(
		newSold.JANCodeList,
		receiver.ToDomain.GetIsVaildJANCodeList(),
		receiver.ToDomain.GetIsVaildLangCodeList(),
		receiver.ToDomain.GetDefaultLangCode(),
	)

	newSold.JANCodeList = janCodeList
	newSold.LanguageCode = languageCode

	// 同じ注文番号が存在するかを確認して、存在する場合はマージ
	if !receiver.ToDomain.MergeWithExistingOrder(newSold) {
		// 存在しない場合は新規に注文リストに追加
		receiver.ToDomain.AddNewSold(newSold)
		receiver.ToDomain.SortOrderList()
	}
}

// DeleteSold ...
func (receiver *useCase) DeleteSold(
	ctx context.Context,
	number int,
) {

	// 削除する注文を見つける
	index := receiver.ToDomain.FindSoldIndex(number)
	if index == -1 {
		// 見つからない場合も問題なし
		return
	}

	// 注文を削除
	receiver.ToDomain.DeleteSoldList(index)

	// 注文リストをソート
	receiver.ToDomain.SortOrderList()

}

// SaveReserving ...
func (receiver *useCase) SaveReserving(
	ctx context.Context,
	newReserving domain.Reserving,
) {

	// 無効なJANcodeを排除
	janCodeList, languageCode := receiver.ToDomain.VerifyJANCodes(
		newReserving.JANCodeList,
		receiver.ToDomain.GetIsVaildJANCodeList(),
		receiver.ToDomain.GetIsVaildLangCodeList(),
		receiver.ToDomain.GetDefaultLangCode(),
	)
	newReserving.JANCodeList = janCodeList
	newReserving.LanguageCode = languageCode

	// 古い queueNo の内容を入れ替え
	// 古い queueNo がなければ新規追加
	if !receiver.ToDomain.UpdateExistingReserving(newReserving) {
		receiver.ToDomain.AddNewReserving(newReserving)
	}
}

// GetPreparingList ...
func (receiver *useCase) GetPreparingList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetPreparingList(
		ctx,
	)
}

// GetCompletedList ...
func (receiver *useCase) GetCompletedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetCompletedList(
		ctx,
	)
}

// GetPassedList ...
func (receiver *useCase) GetPassedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToDomain.GetPassedList(
		ctx,
	)
}

// UpdateSold ...
func (receiver *useCase) UpdateSoldStatus(
	ctx context.Context,
	newSold domain.Sold,
) {
	receiver.ToDomain.UpdateSoldStatus(
		ctx,
		newSold,
	)
}

func (receiver *useCase) DetectSaveJANCodes(
	ctx context.Context,
	number int,
	file *multipart.FileHeader,
) error {
	reserving := domain.Reserving{}

	originalJANCodes, err := reserving.CodeDetector(number, file)
	if err != nil {
		return err
	}

	janCodeList, languageCode := receiver.ToDomain.VerifyJANCodes(
		originalJANCodes,
		receiver.ToDomain.GetIsVaildJANCodeList(),
		receiver.ToDomain.GetIsVaildLangCodeList(),
		receiver.ToDomain.GetDefaultLangCode(),
	)

	reserving.QueueNo = number
	reserving.JANCodeList = janCodeList
	reserving.LanguageCode = languageCode

	// // 古い queueNo の内容を入れ替え
	// // 古い queueNo がなければ新規追加
	if !receiver.ToDomain.UpdateExistingReserving(reserving) {
		receiver.ToDomain.AddNewReserving(reserving)
	}

	return nil
}

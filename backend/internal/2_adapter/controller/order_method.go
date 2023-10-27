package controller

import (
	"context"
	"mime/multipart"

	domain "backend/internal/4_domain"
)

// order -----------------------
// GetOrderList ...
func (receiver *controller) GetOrderList(
	ctx context.Context,
) domain.OrderList {
	return receiver.ToUseCase.GetOrderList(
		ctx,
	)
}

// GetReservingList ...
func (receiver *controller) GetReservingList(
	ctx context.Context,
) domain.ReservingList {
	return receiver.ToUseCase.GetReservingList(
		ctx,
	)
}

// GetReserving ...
func (receiver *controller) GetReserving(
	ctx context.Context,
	number int,
) domain.Reserving {
	return receiver.ToUseCase.GetReserving(
		ctx,
		number,
	)
}

// GetSoldList ...
func (receiver *controller) GetSoldList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetSoldList(
		ctx,
	)
}

// SaveSold ...
func (receiver *controller) SaveSold(
	ctx context.Context,
	newSold domain.Sold,
) {
	receiver.ToUseCase.SaveSold(
		ctx,
		newSold,
	)
}

// SaveReserving ...
func (receiver *controller) SaveReserving(
	ctx context.Context,
	newReserving domain.Reserving,
) {
	receiver.ToUseCase.SaveReserving(
		ctx,
		newReserving,
	)
}

// DeleteSold ...
func (receiver *controller) DeleteSold(
	ctx context.Context,
	number int,
) {
	receiver.ToUseCase.DeleteSold(
		ctx,
		number,
	)
}

// GetPreparingList ...
func (receiver *controller) GetPreparingList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetPreparingList(
		ctx,
	)
}

// GetCompletedList ...
func (receiver *controller) GetCompletedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetCompletedList(
		ctx,
	)
}

// GetPassedList ...
func (receiver *controller) GetPassedList(
	ctx context.Context,
) domain.SoldList {
	return receiver.ToUseCase.GetPassedList(
		ctx,
	)
}

// UpdateSold ...
func (receiver *controller) UpdateSoldStatus(
	ctx context.Context,
	newSold domain.Sold,
) {
	receiver.ToUseCase.UpdateSoldStatus(
		ctx,
		newSold,
	)
}

func (receiver *controller) DetectSaveJANCodes(
	ctx context.Context,
	number int,
	file *multipart.FileHeader,
) error {
	return receiver.ToUseCase.DetectSaveJANCodes(
		ctx,
		number,
		file,
	)
}

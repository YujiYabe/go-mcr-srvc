package domain

import (
	"context"
	"sort"
	"strconv"
)

func (receiver *OrderList) GetReservingList(ctx context.Context) ReservingList {
	return receiver.ReservingList
}

// GetReserving ...
func (receiver *OrderList) GetReserving(
	ctx context.Context,
	number int,
) Reserving {
	targetReserving := Reserving{}
	for _, reserving := range receiver.ReservingList {
		if number == reserving.QueueNo {
			targetReserving = reserving
			break
		}
	}

	return targetReserving
}

func (receiver *OrderList) GetOrderList(ctx context.Context) OrderList {
	return *receiver
}

func (receiver *OrderList) GetSoldList(ctx context.Context) SoldList {
	return receiver.SoldList
}

// SaveSold ...
func (receiver *OrderList) SaveSold(
	ctx context.Context,
	newSold Sold,
) {

}

func (receiver *OrderList) GetPreparingList(ctx context.Context) SoldList {
	return receiver.filterSoldList(StatusPreparing)
}

func (receiver *OrderList) GetCompletedList(ctx context.Context) SoldList {
	return receiver.filterSoldList(StatusCompleted)
}

func (receiver *OrderList) GetPassedList(ctx context.Context) SoldList {
	return receiver.filterSoldList(StatusPassed)
}

func (receiver *OrderList) filterSoldList(statusNum StatusNum) SoldList {
	newSoldList := SoldList{}
	for _, sold := range receiver.SoldList {
		if sold.Status == statusNum {
			newSoldList = append(newSoldList, sold)
		}
	}

	return newSoldList
}

func (receiver *OrderList) PickOutNumber(numberString string) (int, error) {
	return strconv.Atoi(numberString)
}

func PickOutNumber(numberString string) (int, error) {
	return strconv.Atoi(numberString)
}

// UpdateSoldStatus はオーダーリスト内の特定の売れたアイテムのステータスを更新します。
func (receiver *OrderList) UpdateSoldStatus(
	ctx context.Context,
	newSold Sold,
) {
	for index, sold := range receiver.SoldList {
		if sold.SoldNo == newSold.SoldNo {
			receiver.SoldList[index].Status = newSold.Status
			break
		}
	}
}

// FindReservingByNumber は指定されたキュー番号の予約情報を返します。
func (receiver *OrderList) FindReservingByNumber(number int) *Reserving {
	for _, reserving := range receiver.ReservingList {
		if number == reserving.QueueNo {
			return &reserving
		}
	}

	return nil
}

// ResetReservingList は指定された番号の予約リストをリセットします。
func (receiver *OrderList) ResetReservingList(number int) bool {
	for index, reserving := range receiver.ReservingList {
		if number == reserving.QueueNo {
			receiver.ReservingList[index].JANCodeList = []int{}
			return true
		}
	}

	return false
}

// FindPreparingSoldItem はPreparingのステータスを持つSoldアイテムを検索します。
func (receiver *OrderList) FindPreparingSoldItem(targetNumber int, counter *int) interface{} {
	for _, sold := range receiver.SoldList {
		if sold.Status == StatusPreparing {
			*counter++
			if *counter == targetNumber {
				return sold
			}
		}
	}

	return nil
}

// FindReservingItem はReservingListからアイテムを検索します。
func (receiver *OrderList) FindReservingItem(targetNumber int, counter *int) interface{} {
	for _, reserving := range receiver.ReservingList {
		*counter++
		if *counter == targetNumber {
			return reserving
		}
	}

	return nil
}

// FindSoldIndex は指定された注文番号の注文のインデックスを返します。
// 見つからない場合は-1を返します。
func (receiver *OrderList) FindSoldIndex(soldNo int) int {
	for index, sold := range receiver.SoldList {
		if sold.SoldNo == soldNo {
			return index
		}
	}

	return -1
}

func (receiver *OrderList) DeleteSoldList(index int) {
	receiver.SoldList = append(receiver.SoldList[:index], receiver.SoldList[index+1:]...)
}

// MergeWithExistingOrder は新しい売却情報を既存の注文とマージします。
// マージが成功した場合はtrueを、それ以外の場合はfalseを返します。
func (receiver *OrderList) MergeWithExistingOrder(newSold Sold) bool {
	for index, sold := range receiver.SoldList {
		if sold.SoldNo == newSold.SoldNo {
			receiver.SoldList[index].JANCodeList = append(receiver.SoldList[index].JANCodeList, newSold.JANCodeList...)
			return true
		}
	}

	return false
}

func (receiver *OrderList) UpdateExistingReserving(number int, newReserving *Reserving) bool {
	for index, reserving := range receiver.ReservingList {
		if number == reserving.QueueNo {
			receiver.ReservingList[index].JANCodeList = newReserving.JANCodeList
			receiver.ReservingList[index].LanguageCode = newReserving.LanguageCode
			return true
		}
	}
	return false
}

func (receiver *OrderList) AddNewReserving(newReserving *Reserving) {
	// receiver.ReservingList = append(receiver.ReservingList, *newReserving)
	// sort.Slice(receiver.ReservingList, func(i, j int) bool {
	// 	return receiver.ReservingList[i].QueueNo < receiver.ReservingList[j].QueueNo
	// })
}

// 注文リストの更新
func (receiver *OrderList) SortOrderList() {
	// 注文リストをソート
	sort.Slice(receiver.SoldList, func(i, j int) bool {
		return receiver.SoldList[i].SoldNo < receiver.SoldList[j].SoldNo
	})

}

func (receiver *OrderList) AddNewSold(newSold Sold) {
	receiver.SoldList = append(receiver.SoldList, newSold)
}

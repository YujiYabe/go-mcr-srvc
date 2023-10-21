package domain

import (
	"sort"
	"strconv"

	"gocv.io/x/gocv"
)

var orderList *OrderList

type StatusNum int

const (
	StatusPreparing StatusNum = 1
	StatusCompleted StatusNum = 2
	StatusPassed    StatusNum = 3
)

type OrderList struct {
	reservingList []Reserving
	soldList      []Sold
}

type Reserving struct {
	QueueNo      int   `json:"queue_no"`
	LanguageCode int   `json:"language_code"`
	JANCodeList  []int `json:"jan_code_list"`
}

type Sold struct {
	SoldNo       int       `json:"sold_no"`
	LanguageCode int       `json:"language_code"`
	Status       StatusNum `json:"status"`
	JANCodeList  []int     `json:"jan_code_list"`
}

func InitOrder(
	isDemo bool,
) {
	orderList = &OrderList{
		reservingList: []Reserving{},
		soldList:      []Sold{},
	}

	if isDemo {
		orderList.soldList = SoldListDemo
		orderList.reservingList = ReservingListDemo
	}
}

func GetOrderList() *OrderList {
	return orderList
}

func GetReservingList() []Reserving {
	return orderList.reservingList
}

func GetSoldList() []Sold {
	return orderList.soldList
}

func GetPreparingList() []Sold {
	return filterSoldList(StatusPreparing)
}

func GetCompletedList() []Sold {
	return filterSoldList(StatusCompleted)
}

func GetPassedList() []Sold {
	return filterSoldList(StatusPassed)
}

func filterSoldList(statusNum StatusNum) []Sold {
	newSoldList := []Sold{}
	for _, sold := range orderList.soldList {
		if sold.Status == statusNum {
			newSoldList = append(newSoldList, sold)
		}
	}

	return newSoldList
}

func PickOutJANCodes(qcd gocv.QRCodeDetector, img gocv.Mat) []int {
	janCodes := []int{}

	codes := make([]string, 20)
	points := gocv.NewMat()
	qrcodes := make([]gocv.Mat, 20)

	res := qcd.DetectAndDecodeMulti(img, &codes, &points, &qrcodes)
	if !res {
		return janCodes
	}
	splitCodes := []int{}
	for _, code := range codes {
		if code != "" {
			codeInt, _ := strconv.Atoi(code)
			splitCodes = append(splitCodes, codeInt)
		}
	}

	return splitCodes
}

// func (domain OrderList) GetAll() {
// 	fmt.Println("== == == == == == == == == == ")
// 	fmt.Printf("%#v\n", orderList)
// 	fmt.Println("== == == == == == == == == == ")
// }

func PickOutNumber(numberString string) (int, error) {
	return strconv.Atoi(numberString)
}

// UpdateSoldStatus はオーダーリスト内の特定の売れたアイテムのステータスを更新します。
func UpdateSoldStatus(newSold *Sold) {
	for index, sold := range orderList.soldList {
		if sold.SoldNo == newSold.SoldNo {
			orderList.soldList[index].Status = newSold.Status
			break
		}
	}
}

// FindReservingByNumber は指定されたキュー番号の予約情報を返します。
func FindReservingByNumber(number int) *Reserving {
	for _, reserving := range orderList.reservingList {
		if number == reserving.QueueNo {
			return &reserving
		}
	}

	return nil
}

// ResetReservingList は指定された番号の予約リストをリセットします。
func ResetReservingList(number int) bool {
	for index, reserving := range orderList.reservingList {
		if number == reserving.QueueNo {
			orderList.reservingList[index].JANCodeList = []int{}
			return true
		}
	}

	return false
}

// FindPreparingSoldItem はPreparingのステータスを持つSoldアイテムを検索します。
func FindPreparingSoldItem(targetNumber int, counter *int) interface{} {
	for _, sold := range orderList.soldList {
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
func FindReservingItem(targetNumber int, counter *int) interface{} {
	for _, reserving := range orderList.reservingList {
		*counter++
		if *counter == targetNumber {
			return reserving
		}
	}

	return nil
}

// FindSoldIndex は指定された注文番号の注文のインデックスを返します。
// 見つからない場合は-1を返します。
func FindSoldIndex(soldNo int) int {
	for index, sold := range orderList.soldList {
		if sold.SoldNo == soldNo {
			return index
		}
	}

	return -1
}

func DeleteSoldList(index int) {
	orderList.soldList = append(orderList.soldList[:index], orderList.soldList[index+1:]...)
}

// MergeWithExistingOrder は新しい売却情報を既存の注文とマージします。
// マージが成功した場合はtrueを、それ以外の場合はfalseを返します。
func MergeWithExistingOrder(newSold *Sold) bool {
	for index, sold := range orderList.soldList {
		if sold.SoldNo == newSold.SoldNo {
			orderList.soldList[index].JANCodeList = append(orderList.soldList[index].JANCodeList, newSold.JANCodeList...)
			return true
		}
	}

	return false
}

func UpdateExistingReserving(number int, newReserving *Reserving) bool {
	for index, reserving := range orderList.reservingList {
		if number == reserving.QueueNo {
			orderList.reservingList[index].JANCodeList = newReserving.JANCodeList
			orderList.reservingList[index].LanguageCode = newReserving.LanguageCode
			return true
		}
	}
	return false
}

func AddNewReserving(newReserving *Reserving) {
	orderList.reservingList = append(orderList.reservingList, *newReserving)
	sort.Slice(orderList.reservingList, func(i, j int) bool {
		return orderList.reservingList[i].QueueNo < orderList.reservingList[j].QueueNo
	})
}

// 注文リストの更新
func SortOrderList() {
	// 注文リストをソート
	sort.Slice(orderList.soldList, func(i, j int) bool {
		return orderList.soldList[i].SoldNo < orderList.soldList[j].SoldNo
	})
}

func AddNewSold(newSold *Sold) {
	orderList.soldList = append(orderList.soldList, *newSold)
}

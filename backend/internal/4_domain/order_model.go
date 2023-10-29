package domain

type StatusNum int

const (
	StatusPreparing StatusNum = 1
	StatusCompleted StatusNum = 2
	StatusPassed    StatusNum = 3
)

type ViaNum int

const (
	ViaUnknown ViaNum = 0 // 不明
	ViaCasher  ViaNum = 1 // レジ経由
	ViaMobile  ViaNum = 2 // grpc経由でスマホアプリに注文番号を返却する必要あり
)

type (
	OrderList struct {
		ReservingList
		SoldList
	}

	ReservingList []Reserving
	Reserving     struct {
		QueueNo      int   `json:"queue_no"`
		LanguageCode int   `json:"language_code"`
		JANCodeList  []int `json:"jan_code_list"`
	}

	SoldList []Sold
	Sold     struct {
		SoldNo       int       `json:"sold_no"`
		LanguageCode int       `json:"language_code"`
		Via          ViaNum    `json:"via"`
		Status       StatusNum `json:"status"`
		JANCodeList  []int     `json:"jan_code_list"`
	}
)

func NewOrderList(
	isDemo bool,
) OrderList {
	orderList := OrderList{
		ReservingList: NewReservingList(),
		SoldList:      NewSoldList(),
	}
	if isDemo {
		orderList.SoldList = SoldListDemo
		orderList.ReservingList = ReservingListDemo
	}

	return orderList

}

func NewReservingList() ReservingList {
	return ReservingList{}
}

func NewSoldList() SoldList {
	return SoldList{}
}

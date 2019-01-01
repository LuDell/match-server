package model

type Trade struct {
	Symbol string
	ContractId uint
	Price float64
	Volume uint
	BidId uint
	AskId uint
	TrendSide string
	BidUid uint
	AskUid uint
	BidFromStatus uint8
	BidToStatus uint8
	AskToStatus uint8
	Ctime int
	Token string
}

func (t *Trade)CheckTrade(bidOrder Order,askOrder Order) *bool {

}

func (t *Trade)chargeFee(order Order) (uint,error) {
	if t.TrendSide == order.Side {

	}
}
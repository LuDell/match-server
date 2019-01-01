package model

type Order struct {
	Id  uint
	Uid uint
	ContractId uint
	UnitQuantity float64
	Side string
	Action string
	Price float64
	Volume uint
	FeeRateMaker float64
	FeeRateTaker float64
	DealVolume uint
	LeverageLevel uint
	AvgPrice float64
	RealizedAmount float64
	Status uint8
	Type uint8
	Ctime int64
	Mtime int64
	Source uint8
	OperType uint8
	OriginOid uint
}

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

type Position struct {

	Id uint

	/**
	 * 用户ID
	 */
	Uid uint

	/**
	 * 合约ID
	 */
	ContractId uint

	/**
	 * 仓位类型：1 全仓，2 逐仓
	 */
	CopType uint8

	/**
	 * 买卖方向：BUY 多，SELL 空
	 */
	Side string

	/**
	 * 持仓保证金(起始保证金+变动)
	 */
	HoldAmount float64

	/**
	 * 冻结平仓手续费
	 */
	CloseFeeAmount float64

	/**
	 * 开仓均价
	 */
	AvgPrice float64

	/**
	 * 持仓总量，单位：张
	 */
	Volume uint

	/**
	 * 杠杆倍数
	 */
	LeverageLevel uint

	/**
	 * 被强制平仓的用户爆仓单ID
	 */
	OriginOid uint

	/**
	 * 用户持仓冻结状态：0 正常，1爆仓冻结，2 强减冻结，3 交割冻结
	 */
	FreezeLock uint8

	/**
	 * 历史已实现盈亏合计金额
	 */
	HistoryRealizedAmount float64

	/**
	 * 用户分组
	 */
	UserGroup uint8

	/**
	 * 创建时间
	 */
	Ctime int64

	/**
	 * 更新时间
	 */
	Mtime int64
}

type Transaction struct {
	Id uint
	FromUid uint
	FromType int
	FromBalance float64
	ToUid uint
	ToType int
	ToBalance float64
	Amount float64
	Meta string
	Scene string
	RefType string
	RefId uint
	Op_uid uint
	Op_ip string
	Ctime int64
	Mtime int64
	Fingerprint string
}

type Trade struct {

}

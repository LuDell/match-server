package model

type CoPosition struct {

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

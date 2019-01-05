package model

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
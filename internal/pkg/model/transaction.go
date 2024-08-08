package model

import (
	"time"
)

type Transaction struct {
	TransactionId   uint `gorm:"primaryKey;autoIncrement"`
	AccountId       uint
	OperationTypeId uint
	Amount          float64
	Balance         float64
	EventDate       time.Time
}

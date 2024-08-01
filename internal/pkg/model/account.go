package model

type Account struct {
	AccountId      uint   `gorm:"primaryKey;autoIncrement" json:"account_id,omitempty"`
	DocumentNumber string `gorm:"uniqueIndex" json:"document_number"`
}

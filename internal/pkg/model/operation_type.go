package model

import "errors"

type OperationType int

const (
	NORMAL_PURCHASE OperationType = iota + 1
	PURCHASE_WITH_INSTALLMENTS
	WITHDRAWL
	CREDIT_VOUCHER
)

var PurchaseTypes = map[OperationType]bool{
	NORMAL_PURCHASE:            true,
	PURCHASE_WITH_INSTALLMENTS: true,
	WITHDRAWL:                  true,
}

func GetOperationType(opType int) (OperationType, error) {
	switch opType {
	case int(NORMAL_PURCHASE), int(PURCHASE_WITH_INSTALLMENTS), int(WITHDRAWL), int(CREDIT_VOUCHER):
		return OperationType(opType), nil
	}
	return 0, errors.New("invalid Operation Type")
}

func (tt OperationType) String() string {
	return [...]string{"", "Normal Purchase", "Purchase with Installments", "Withdrawal", "Credit Voucher"}[tt]
}

func (tt OperationType) IsPurchaseType() bool {
	_, ok := PurchaseTypes[tt]
	return ok
}

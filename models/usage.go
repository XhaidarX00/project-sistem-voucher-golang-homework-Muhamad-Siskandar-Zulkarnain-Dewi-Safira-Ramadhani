package models

import "time"

type Usage struct {
	Base
	UserID            int       `json:"user_id" gorm:"type:integer"`
	VoucherID         int       `gorm:"type:integer;index"` // Foreign key to Voucher
	VoucherCode       string    `gorm:"type:varchar(255)" json:"voucher_code"`
	UsageDate         time.Time `gorm:"type:timestamp"`
	TransactionAmount float64   `gorm:"type:decimal(10,2)"`
	BenefitAmount     float64   `gorm:"type:decimal(10,2)"`

	// Relationships
	Voucher Voucher `gorm:"foreignKey:VoucherID"`
}

type UsageDTO struct {
	UserID       int        `json:"user_id"`
	VoucherInput VoucherDTO `json:"voucher_input"`
}

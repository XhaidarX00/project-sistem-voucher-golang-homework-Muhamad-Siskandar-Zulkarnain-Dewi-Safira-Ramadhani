package models

import "time"

type Voucher struct {
	ID              int       `gorm:"primaryKey;autoIncrement" json:"id"`
	VoucherName     string    `gorm:"type:varchar(255)" json:"voucher_name"`
	VoucherCode     string    `gorm:"type:varchar(255);uniqueIndex" json:"voucher_code"`
	VoucherType     string    `gorm:"type:voucher_type" json:"voucher_type"`
	Description     string    `gorm:"type:text" json:"description"`
	VoucherCategory string    `gorm:"type:voucher_category" json:"voucher_category"`
	DiscountAmount  float64   `gorm:"type:decimal(10,2)" json:"discount_amount"`
	MinPurchase     float64   `gorm:"type:decimal(10,2)" json:"min_purchase"`
	PaymentMethod   string    `gorm:"type:varchar(255)" json:"payment_method"`
	StartDate       time.Time `gorm:"type:date" json:"start_date"`
	EndDate         time.Time `gorm:"type:date" json:"end_date"`
	ApplicableAreas string    `gorm:"type:jsonb" json:"applicable_areas"`

	// Relationships
	Redeems []Redeem `gorm:"foreignKey:VoucherID"`
	// Usages  []Usage  `gorm:"foreignKey:VoucherID"`
}

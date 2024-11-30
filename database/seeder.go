package database

import (
	"project-voucher-team3/models"
	"time"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	err := voucherSeeder(db)
	if err != nil {
		return err
	}
	err = redeemSeeder(db)
	if err != nil {
		return err
	}
	return nil
}

func voucherSeeder(db *gorm.DB) error {
	// Seed Vouchers
	vouchers := []models.Voucher{
		{
			VoucherName:     "Free Shipping Java",
			VoucherCode:     "FREESHIPJAVA",
			VoucherType:     "ecommerce",
			Description:     "Free shipping for orders in Java island.",
			VoucherCategory: "free_shipping",
			DiscountAmount:  0,
			MinPurchase:     50000,
			PaymentMethod:   "credit_card",
			StartDate:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			ApplicableAreas: `["Jakarta", "Surabaya", "Bandung"]`,
		},
		{
			VoucherName:     "10% Off Electronics",
			VoucherCode:     "ELEC10",
			VoucherType:     "ecommerce",
			Description:     "10% discount on electronics above Rp500,000.",
			VoucherCategory: "discount",
			DiscountAmount:  10,
			MinPurchase:     500000,
			PaymentMethod:   "debit_card",
			StartDate:       time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2024, 11, 30, 23, 59, 59, 0, time.UTC),
			ApplicableAreas: `["Jakarta", "Medan"]`,
		},
	}
	if err := db.Create(&vouchers).Error; err != nil {
		return err
	}
	return nil
}

func redeemSeeder(db *gorm.DB) error {
	// Fetch Voucher IDs
	var voucher1, voucher2 models.Voucher
	if err := db.First(&voucher1, "voucher_code = ?", "ELEC10").Error; err != nil {
		return err
	}
	if err := db.First(&voucher2, "voucher_code = ?", "FREESHIPJAVA").Error; err != nil {
		return err
	}

	// Seed Redeems
	redeems := []models.Redeem{
		{
			UserID:      1,
			VoucherID:   voucher1.ID,
			VoucherCode: "ELEC10",
			RedeemDate:  time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
		},
		{
			UserID:      2,
			VoucherID:   voucher2.ID,
			VoucherCode: "FREESHIPJAVA",
			RedeemDate:  time.Date(2024, 4, 10, 12, 0, 0, 0, time.UTC),
		},
	}
	if err := db.Create(&redeems).Error; err != nil {
		return err
	}
	return nil
}

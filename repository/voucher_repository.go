package repository

import (
	"project-voucher-team3/models"

	"gorm.io/gorm"
)

type VoucherRepository struct {
	DB *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{db}
}

func (repo *VoucherRepository) GetUserVoucher(voucherFilter models.Voucher) (models.Voucher, error) {
	var voucher models.Voucher
	err := repo.DB.Where("voucher_code = ? AND voucher_type = ?", voucherFilter.VoucherCode, voucherFilter.VoucherType).First(&voucher).Error
	return voucher, err
}

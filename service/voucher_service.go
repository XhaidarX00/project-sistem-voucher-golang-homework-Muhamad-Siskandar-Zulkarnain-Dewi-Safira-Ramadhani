package service

import (
	"project-voucher-team3/models"
	"project-voucher-team3/repository"

	"gorm.io/gorm"
)

type VoucherService interface {
	GetUserVoucher(voucherFilter models.Voucher) (models.Voucher, error)
}

type voucherService struct {
	Repo repository.Repository
}

func NewVoucherService(db *gorm.DB) VoucherService {
	return &voucherService{Repo: repository.NewRepository(db)}
}

func (s *voucherService) GetUserVoucher(voucherFilter models.Voucher) (models.Voucher, error) {
	return s.Repo.Voucher.GetUserVoucher(voucherFilter)
}

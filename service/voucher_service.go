package service

import (
	"project-voucher-team3/models"
	"project-voucher-team3/repository"
	"project-voucher-team3/utils"
	"strings"
	"time"
)

type VoucherService interface {
	ValidateVoucher(voucherInput models.VoucherDTO) (*models.ValidateVoucherResponse, error)
}

type voucherService struct {
	Repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo}
}

func (s *voucherService) ValidateVoucher(voucherInput models.VoucherDTO) (*models.ValidateVoucherResponse, error) {
	voucher, err := s.Repo.GetVoucherByCode(voucherInput.VoucherCode)
	if err != nil {
		return nil, err
	}
	customDateFormat := "2006-01-02"
	str := strings.Trim(string(voucherInput.TransactionDate), `"`)
	parsedTime, err := time.Parse(customDateFormat, str)
	if err != nil {
		return nil, err
	}
	voucherInput.FormatedTransactionDate = parsedTime
	result, err := utils.ValidateVoucher(voucherInput, voucher)
	return &result, err
}

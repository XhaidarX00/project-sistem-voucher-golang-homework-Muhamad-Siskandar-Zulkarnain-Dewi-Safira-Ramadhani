package service

import (
	"encoding/json"
	"errors"
	"math"
	"project-voucher-team3/models"
	"project-voucher-team3/repository"
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
	var result models.ValidateVoucherResponse

	// Fetch the voucher by code
	voucher, err := s.Repo.GetVoucherByCode(voucherInput.VoucherCode)
	if err != nil {
		return nil, err
	}
	if voucher.ID == 0 {
		return nil, errors.New("voucher not found")
	}

	// Check if the voucher is active
	if !isVoucherActive(voucher, voucherInput.TransactionDate.ToTime()) {
		return nil, errors.New("voucher is already out of date")
	}
	result.VoucherStatus = "Active"

	// Validate total transaction
	if voucherInput.TotalTransaction < voucher.MinPurchase {
		return nil, errors.New("your total transaction must be greater than the minimum purchase")
	}
	result.TotalTransaction = voucherInput.TotalTransaction
	result.TotalShippingCost = voucherInput.TotalShippingCost

	// Validate payment method
	if voucher.PaymentMethod != "" && voucherInput.PaymentMethod != voucher.PaymentMethod {
		return nil, errors.New("payment method does not match")
	}

	// Validate area
	if voucher.ApplicableAreas != "" && !isValidArea(voucher.ApplicableAreas, voucherInput.Area) {
		return nil, errors.New("area is not valid")
	}

	// Calculate benefit amount
	result.BenefitAmount = calculateBenefit(voucher, voucherInput)

	return &result, nil
}

// Helper to check if the voucher is active
func isVoucherActive(voucher models.Voucher, transactionDate time.Time) bool {
	return voucher.StartDate.Before(transactionDate) && voucher.EndDate.After(transactionDate)
}

// Helper to validate the area
func isValidArea(applicableAreasJSON, inputArea string) bool {
	var areas []string
	if err := json.Unmarshal([]byte(applicableAreasJSON), &areas); err != nil {
		return false
	}
	for _, area := range areas {
		if area == inputArea {
			return true
		}
	}
	return false
}

// Helper to calculate the benefit amount
func calculateBenefit(voucher models.Voucher, voucherInput models.VoucherDTO) float64 {
	switch voucher.VoucherCategory {
	case "discount":
		return math.Round(voucherInput.TotalTransaction*((100.00-voucher.DiscountAmount)/100.00)*100) / 100
	case "free_shipping":
		return math.Round(voucherInput.TotalShippingCost*((100.00-voucher.DiscountAmount)/100.00)*100) / 100
	default:
		return 0
	}
}

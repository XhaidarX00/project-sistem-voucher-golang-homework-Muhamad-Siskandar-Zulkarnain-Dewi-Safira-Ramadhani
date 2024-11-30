package service

import (
	"project-voucher-team3/models"
	"project-voucher-team3/repository"
)

type RedeemService interface {
	GetAllUserRedeems(userID int, voucherFilter models.Voucher) ([]models.Redeem, error)
}

type redeemService struct {
	Repo repository.ReedemRepository
}

func NewRedeemService(repo repository.ReedemRepository) RedeemService {
	return &redeemService{repo}
}

func (s *redeemService) GetAllUserRedeems(userID int, voucherFilter models.Voucher) ([]models.Redeem, error) {
	return s.Repo.GetUserRedeem(userID, voucherFilter)
}

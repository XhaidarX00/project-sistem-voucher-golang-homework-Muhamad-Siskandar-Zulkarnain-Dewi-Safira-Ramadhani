package service

import (
	"project-voucher-team3/repository"
)

type Service struct {
	User    UserService
	Reedem  RedeemService
	Voucher VoucherService
}

func NewService(repo repository.Repository) Service {
	return Service{
		User:    NewUserService(repo.User),
		Reedem:  NewRedeemService(repo.Reedem),
		Voucher: NewVoucherService(repo.Voucher),
	}
}

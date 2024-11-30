package service

import (
	"project-voucher-team3/repository"

	"gorm.io/gorm"
)

type Service struct {
	User   UserService
	Reedem RedeemService
}

func NewService(repo repository.Repository, db *gorm.DB) Service {
	return Service{
		User:   NewUserService(repo.User),
		Reedem: NewRedeemService(repo.Reedem),
	}
}

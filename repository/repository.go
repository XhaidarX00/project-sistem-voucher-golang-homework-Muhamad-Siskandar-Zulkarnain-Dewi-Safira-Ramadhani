package repository

import "gorm.io/gorm"

type Repository struct {
	User    UserRepository
	Voucher VoucherRepository
	Reedem  ReedemRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		User:    *NewUserRepository(db),
		Voucher: *NewVoucherRepository(db),
		Reedem:  *NewReedemRepository(db),
	}
}

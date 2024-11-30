package repository

import (
	"project-voucher-team3/models"
	"time"

	"gorm.io/gorm"
)

type ReedemRepository struct {
	DB *gorm.DB
}

func NewReedemRepository(db *gorm.DB) *ReedemRepository {
	return &ReedemRepository{db}
}

func (repo *ReedemRepository) GetUserRedeem(userID int, voucherFilter models.Voucher) ([]models.Redeem, error) {
	var redeems []models.Redeem

	err := repo.DB.Preload("Voucher", "voucher_type = ?", voucherFilter.VoucherType).
		Where("user_id = ?", userID).
		Find(&redeems).Error
	if err != nil {
		return nil, err
	}

	now := time.Now()
	activeRedeems := []models.Redeem{}
	for _, redeem := range redeems {
		if redeem.Voucher.ID != 0 && redeem.Voucher.StartDate.Before(now) && redeem.Voucher.EndDate.After(now) {
			activeRedeems = append(activeRedeems, redeem)
		}
	}

	return activeRedeems, nil
}

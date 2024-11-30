package database

import (
	"project-voucher-team3/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.Exec(`
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'voucher_type') THEN
				CREATE TYPE voucher_type AS ENUM ('ecommerce', 'redeem_point');
			END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'voucher_category') THEN
				CREATE TYPE voucher_category AS ENUM ('free_shipping', 'discount');
			END IF;
		END $$;
	`).Error
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.Voucher{},
		&models.Redeem{},
	)
	return err
}

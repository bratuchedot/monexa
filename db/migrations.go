package db

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"monexa/models"
	"os"
	"time"
)

func createUniqueEmailIndex(tx *gorm.DB) error {
	query := `
		CREATE UNIQUE INDEX IF NOT EXISTS unique_email_active_users 
		ON users (email) 
		WHERE deleted_at IS NULL;
	`
	return tx.Exec(query).Error
}

func dropUniqueEmailIndex(tx *gorm.DB) error {
	query := `
		DROP INDEX IF EXISTS unique_email_active_users;
	`
	return tx.Exec(query).Error
}

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20241214094152_create_users_table",
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&models.User{})
				if err != nil {
					return err
				}
				return createUniqueEmailIndex(tx)
			},
			Rollback: func(tx *gorm.DB) error {
				err := dropUniqueEmailIndex(tx)
				if err != nil {
					return err
				}
				return tx.Migrator().DropTable("users")
			},
		},
		{
			ID: "20241214102308_create_settings_table",
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&models.Setting{})
				if err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("settings")
			},
		},
		{
			ID: "20241214102915_create_categories_table",
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&models.Category{})
				if err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("categories")
			},
		},
		{
			ID: "20241214102929_create_payment_methods_table",
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&models.PaymentMethod{})
				if err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("payment_methods")
			},
		},
		{
			ID: "20241214103104_create_records_table",
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&models.Record{})
				if err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("records")
			},
		},
		{
			ID: "20241214104212_add_test_user",
			Migrate: func(tx *gorm.DB) error {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
				if err != nil {
					return err
				}

				testUser := map[string]interface{}{
					"email":      "test@monexa.com",
					"password":   string(hashedPassword),
					"name":       "Test User",
					"created_at": time.Now(),
				}
				if err := tx.Table("users").Create(testUser).Error; err != nil {
					return err
				}

				var userID uint
				if err := tx.Table("users").Select("id").Where("email = ?", "test@monexa.com").Scan(&userID).Error; err != nil {
					return err
				}

				setting := map[string]interface{}{
					"user_id":  userID,
					"language": "EN",
					"currency": "MKD",
				}
				if err := tx.Table("settings").Create(setting).Error; err != nil {
					return err
				}

				paymentMethods := []map[string]interface{}{
					{
						"user_id": userID,
						"name":    "Cash",
					},
					{
						"user_id": userID,
						"name":    "Card",
					},
				}
				if err := tx.Table("payment_methods").Create(paymentMethods).Error; err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Exec(`
					DELETE FROM payment_methods WHERE user_id = (SELECT id FROM users WHERE email = 'test@monexa.com');
					DELETE FROM settings WHERE user_id = (SELECT id FROM users WHERE email = 'test@monexa.com');
					DELETE FROM users WHERE email = 'test@monexa.com';
				`).Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		fmt.Fprint(os.Stderr, "⛔ ️Exit!!! Cannot apply migrations\n")
		os.Exit(1)
	}
}

package jobs

import (
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"payment-queue/db"
	"payment-queue/models"
)

func ProcessQueue() {
	for {
		var payment models.Payment
		tx := db.DB.Begin()
		if tx.Error != nil {
			log.Error().Err(tx.Error).Msg("Failed to start transaction in queue")
			time.Sleep(1 * time.Second)
			continue
		}

		// lock the oldest pending payment
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("status = ?", "pending").
			Order("created_at ASC").
			First(&payment).Error

		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			log.Error().Err(err).Msg("Failed to fetch payment in queue")
			tx.Rollback()
			time.Sleep(1 * time.Second)
			continue
		}

		log.Info().
			Uint("id", payment.ID).
			Str("user_id", payment.UserID).
			Float64("amount", payment.Amount).
			Str("product_id", payment.ProductID).
			Msg("Processing payment")

		// simulate processing (replace with Stripe or other gateway)
		time.Sleep(2 * time.Second)

		payment.Status = "processed" // Or "failed" on error
		if err := tx.Save(&payment).Error; err != nil {
			log.Error().Err(err).Uint("id", payment.ID).Msg("Failed to update payment")
			tx.Rollback()
			continue
		}

		log.Info().Uint("id", payment.ID).Msg("Payment processed")
		tx.Commit()
	}
}

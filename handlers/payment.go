package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"payment-queue/db"
	"payment-queue/models"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: zerolog.NewConsoleWriter()})
}

// EnqueuePayment godoc
// @Summary Enqueue a new payment
// @Description Add a payment to the processing queue
// @Tags payments
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Bearer token"
// @Param payment body models.Payment true "Payment details"
// @Success 200 {object} models.Payment
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 429 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments [post]
func EnqueuePayment(ctx *gin.Context) {
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		log.Error().Err(err).Msg("Invalid JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&payment); err != nil {
		log.Error().Err(err).Msg("Validation failed")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.Status = "pending"
	payment.Timestamp = time.Now()
	if err := db.DB.Create(&payment).Error; err != nil {
		log.Error().Err(err).Msg("Failed to enqueue payment")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enqueue payment"})
		return
	}

	log.Info().
		Uint("id", payment.ID).
		Str("user_id", payment.UserID).
		Float64("amount", payment.Amount).
		Str("product_id", payment.ProductID).
		Msg("Payment enqueued")
	ctx.JSON(http.StatusOK, payment)
}

// ListPayments godoc
// @Summary List all payments
// @Description Retrieve a list of all payments in the queue
// @Tags payments
// @Produce json
// @Param Authorization header string true "JWT Bearer token"
// @Success 200 {array} models.Payment
// @Failure 401 {object} map[string]string
// @Failure 429 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments [get]
func ListPayments(ctx *gin.Context) {
	var payments []models.Payment
	if err := db.DB.Find(&payments).Error; err != nil {
		log.Error().Err(err).Msg("Failed to list payments")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list payments"})
		return
	}

	log.Info().Int("count", len(payments)).Msg("Listed payments")
	ctx.JSON(http.StatusOK, payments)
}

// GetPayment godoc
// @Summary Get a payment by ID
// @Description Retrieve details of a specific payment
// @Tags payments
// @Produce json
// @Param Authorization header string true "JWT Bearer token"
// @Param id path string true "Payment ID"
// @Success 200 {object} models.Payment
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 429 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/{id} [get]
func GetPayment(ctx *gin.Context) {
	id := ctx.Param("id")
	var payment models.Payment
	if err := db.DB.First(&payment, id).Error; err != nil {
		log.Error().Err(err).Str("id", id).Msg("Payment not found")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	log.Info().Uint("id", payment.ID).Msg("Retrieved payment")
	ctx.JSON(http.StatusOK, payment)
}

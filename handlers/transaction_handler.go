package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/chinotanismail/internal-transfer-system/config"
	"github.com/chinotanismail/internal-transfer-system/models"
)

func CreateTransaction(c *gin.Context) {
	var input struct {
		SourceAccountID      uint   `json:"source_account_id"`
		DestinationAccountID uint   `json:"destination_account_id"`
		Amount               string `json:"amount"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	amount, err := strconv.ParseFloat(input.Amount, 64)
	if err != nil || amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	var status string = "success"
	var transaction models.Transaction

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		var source, dest models.Account

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&source, input.SourceAccountID).Error; err != nil {
			status = "failed"
			return err
		}
		if source.Balance < amount {
			status = "failed"
			return errors.New("insufficient funds")
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&dest, input.DestinationAccountID).Error; err != nil {
			status = "failed"
			return err
		}

		source.Balance -= amount
		dest.Balance += amount

		if err := tx.Save(&source).Error; err != nil {
			status = "failed"
			return err
		}
		if err := tx.Save(&dest).Error; err != nil {
			status = "failed"
			return err
		}

		transaction = models.Transaction{
			SourceAccountID:      source.AccountID,
			DestinationAccountID: dest.AccountID,
			Amount:               amount,
			Status:               status,
		}
		return nil
	})

	if err != nil {
		// Log failed transaction with error reason
		transaction = models.Transaction{
			SourceAccountID:      input.SourceAccountID,
			DestinationAccountID: input.DestinationAccountID,
			Amount:               amount,
			Status:               "failed",
			ErrorReason:          err.Error(),
		}
		config.DB.Create(&transaction)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log successful transaction
	transaction.Status = "success"
	config.DB.Create(&transaction)
	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction created successfully",
		"transaction": transaction,
	})
}

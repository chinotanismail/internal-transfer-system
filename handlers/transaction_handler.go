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

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		var source, dest models.Account

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&source, input.SourceAccountID).Error; err != nil {
			return err
		}
		if source.Balance < amount {
			return errors.New("insufficient funds")
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&dest, input.DestinationAccountID).Error; err != nil {
			return err
		}

		source.Balance -= amount
		dest.Balance += amount

		if err := tx.Save(&source).Error; err != nil {
			return err
		}
		if err := tx.Save(&dest).Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			SourceAccountID:      source.AccountID,
			DestinationAccountID: dest.AccountID,
			Amount:               amount,
		}
		return tx.Create(&transaction).Error
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

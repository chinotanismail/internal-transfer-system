package handlers

import (
	"net/http"
	"strconv"

	"github.com/chinotanismail/internal-transfer-system/config"
	"github.com/chinotanismail/internal-transfer-system/models"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var input struct {
		AccountID      uint   `json:"account_id"`
		InitialBalance string `json:"initial_balance"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	balance, err := strconv.ParseFloat(input.InitialBalance, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid balance format"})
		return
	}

	account := models.Account{
		AccountID: input.AccountID,
		Balance:   balance,
	}

	if err := config.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	c.Status(http.StatusOK)
}

func GetAccount(c *gin.Context) {
	id := c.Param("account_id")

	var account models.Account
	if err := config.DB.First(&account, "account_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account_id": account.AccountID, "balance": strconv.FormatFloat(account.Balance, 'f', 5, 64)})
}

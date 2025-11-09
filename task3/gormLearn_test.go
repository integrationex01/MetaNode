package task3_test

import (
	"MetaNode/task3"
	"testing"

	"gorm.io/gorm"
)

func TestTransactionGormOperations(t *testing.T) {
	db := task3.DbClient()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db            *gorm.DB
		fromAccountID uint
		toAccountID   uint
		amount        float64
	}{
		// TODO: Add test cases.
		{
			name:          "Test Insufficient Balance",
			db:            db,
			fromAccountID: uint(1),
			toAccountID:   uint(2),
			amount:        100.0,
		},
		{
			name:          "Test Successful Transaction",
			db:            db,
			fromAccountID: uint(1),
			toAccountID:   uint(2),
			amount:        50.0,
		},
	}
	db.AutoMigrate(&task3.Account{}, &task3.Transaction{})
	db.Create(&task3.Account{ID: 1, Balance: 50.0})
	db.Create(&task3.Account{ID: 2, Balance: 100.0})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := task3.TransactionGormOperations(tt.db, tt.fromAccountID, tt.toAccountID, tt.amount)
			if gotErr != nil {
				t.Logf("fromAccount: %v, toAccount: %v", db.Where("id = ?", tt.fromAccountID).First(&task3.Account{}), db.Where("id = ?", tt.toAccountID).First(&task3.Account{}))
			} else {
				t.Logf("Transaction successful: trans=%v", db.Last(&task3.Transaction{}))
			}
		})
	}
}

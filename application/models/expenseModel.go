package models

import (
	"time"
)

type User struct {
	Id            uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CreatedAt     time.Time `json:"createdAt"`
	ExpenseBudget uint32    `json:"expenseBudget"`
	Name          string    `json:"name" gorm:"size:40;default:'user'"`
	PhoneNumber   string    `json:"phoneNumber" gorm:"size:15;unique"`
	CountryCode   string    `json:"countryCode" gorm:"size:5"`
	EmailId       string    `json:"emailId" gorm:"size:50;unique"`
}

type Expense struct {
	Id             uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserID         uint      `gorm:"not null;foreignKey:UserId"`
	Title          string    `json:"title" gorm:"size:25;not null"`
	Description    string    `json:"description" gorm:"size:250;not null"`
	CreatedAt      time.Time `json:"createdAt"`
	LastUpdatedAt  time.Time `json:"lastUpdatedAt"`
	Category       string    `json:"category" gorm:"size:20"`
	Amount         int32     `json:"amount"`
	IsRecurring    bool      `json:"isRecurring"`
	Currency       string    `json:"currency"`
	RecurAfterDays uint16    `json:"recurAfterDays"`
}

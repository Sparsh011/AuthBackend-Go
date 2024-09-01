package authpkg

import (
	"time"
)

type User struct {
	Id            uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	CreatedAt     time.Time `json:"createdAt"`
	ExpenseBudget uint32    `json:"expenseBudget"`
	Name          string    `json:"name" gorm:"size:40;default:'user'"`
	PhoneNumber   string    `json:"phoneNumber" gorm:"size:15;unique"`
	EmailId       string    `json:"emailId" gorm:"size:50;unique"`
}

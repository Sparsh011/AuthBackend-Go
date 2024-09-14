package authpkg

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID      `json:"id" gorm:"unique;primaryKey"`
	VerificationTime time.Time      `json:"verificationTime"`
	ExpenseBudget    int32          `json:"expenseBudget"`
	Name             string         `json:"name" gorm:"size:40;default:'user'"`
	PhoneNumber      sql.NullString `json:"phoneNumber" gorm:"size:18;unique"`
	EmailId          sql.NullString `json:"emailId" gorm:"size:50;unique"`
	ProfileUri       string         `json:"profileUri"`
}

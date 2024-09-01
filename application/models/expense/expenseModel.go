package expense

import "time"

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

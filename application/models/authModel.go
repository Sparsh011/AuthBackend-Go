package models

import "gorm.io/gorm"

type SendOtpRequest struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	OtpLength   int    `json:"otpLength,omitempty" validate:"required"`
	Expiry      int    `json:"expiry,omitempty" validate:"required"`
	Channel     string `json:"channel,omitempty" validate:"required"`
}

type VerifyOtpRequest struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	Otp         string `json:"otp,omitempty" validate:"required"`
	OrderId     string `json:"orderId,omitempty" validate:"required"`
}

type ResendOtpRequest struct {
	OrderId string `json:"orderId,omitempty" validate:"required"`
}

type User struct {
	Name               string `json:"name"`
	PhoneNumber        string `json:"phoneNumber"`
	CountryCode        string `json:"countryCode"`
	AccessToken        string `json:"accessToken"`
	RefreshToken       string `json:"refreshToken"`
	LastLoginAt        int64  `json:"lastLoginAt"`
	RefreshTokenExpiry int64  `json:"refreshTokenExpiry"`
}

// DBUser struct for the GORM model
type DBUser struct {
	gorm.Model
	Name               string `gorm:"size:20"`
	PhoneNumber        string `gorm:"unique;size:15"`
	CountryCode        string `gorm:"size:5"`
	RefreshToken       string `gorm:"type:text"`
	LastLoginAt        int64  `gorm:"autoUpdateTime"`
	RefreshTokenExpiry int64  `gorm:"not null"`
}

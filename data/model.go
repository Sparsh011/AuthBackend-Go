package data

type OtpData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

type VerifyOtp struct {
	User *OtpData `json:"user,omitempty" validate:"required"`
	Otp  string   `json:"otp,omitempty" validate:"required"`
}

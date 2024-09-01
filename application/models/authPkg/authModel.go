package authpkg

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

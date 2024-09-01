package authpkg

type AccessTokenRequest struct {
	RefreshToken string `json:"refresh"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access"`
}

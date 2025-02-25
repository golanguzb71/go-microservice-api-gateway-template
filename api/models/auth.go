package models

import (
	"gitlab.udevs.io/eld/eld_go_api_gateway/genproto/company_service"
)

type LoginRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PlatformKey string `json:"platform_key"`
	DeviceName  string `json:"device_name"`
}

type DriverLoginRequest struct {
	Username    string                        `json:"username"`
	Password    string                        `json:"password"`
	PlatformKey string                        `json:"platform_key"`
	DeviceName  string                        `json:"device_name"`
	Force       bool                          `json:"force"`
	DeviceInfo  *company_service.DriverDevice `json:"device_info"`
}

type ResetPasswordRequest struct {
	TokenKey string `json:"token"`
	Password string `json:"password"`
}

type ResetTokenRequest struct {
	TokenKey string `json:"token"`
}

type ResetTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type TokenGenerationResponse struct {
	CompanyUserToken string `json:"company_user_token"`
}

type VerifyEmail struct {
	Verify bool `json:"verify"`
}

type GetMenu struct {
	Token string `json:"token"`
}

type Menu struct {
	Id        string `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	ParentId  string `json:"parent_id"`
	CanCreate bool   `json:"can_create"`
	CanUpdate bool   `json:"can_update"`
	CanGet    bool   `json:"can_get"`
	CanDelete bool   `json:"can_delete"`
	Icon      string `json:"icon"`
}

type UploadResponse struct {
	OriginalURL string `json:"filename,omitempty"`
	Key         string `json:"key"`
}

type SignUpLink struct {
	Link string `json:"link"`
}

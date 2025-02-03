package dto

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
}

type SocialMediaResponse struct {
	ID             uint         `json:"id"`
	Name           string       `json:"name"`
	SocialMediaURL string       `json:"social_media_url"`
	UserID         uint         `json:"user_id"`
	User           UserResponse `json:"user,omitempty"`
}

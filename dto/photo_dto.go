package dto

type PhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
}

type PhotoResponse struct {
	ID       uint         `json:"id"`
	Title    string       `json:"title"`
	Caption  string       `json:"caption"`
	PhotoURL string       `json:"photo_url"`
	UserID   uint         `json:"user_id"`
	User     UserResponse `json:"user,omitempty"`
}

package dto

type CommentRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoID uint   `json:"photo_id" validate:"required"`
}

type CommentResponse struct {
	ID      uint          `json:"id"`
	Message string        `json:"message"`
	PhotoID uint          `json:"photo_id"`
	UserID  uint          `json:"user_id"`
	User    UserResponse  `json:"user,omitempty"`
	Photo   PhotoResponse `json:"photo,omitempty"`
}

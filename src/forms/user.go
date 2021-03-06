package forms

type UserSignup struct {
	Name     string `json:"name" binding:"required"`
	BirthDay string `json:"birthday" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type UserSignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

package auth

type RegisterUserInput struct {
	Name		string `json:"name" binding:"required"`
	Username	string `json:"username" binding:"required"`
}
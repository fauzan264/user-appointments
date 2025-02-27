package auth

type RegisterUserInput struct {
	Name				string `json:"name" binding:"required"`
	Username			string `json:"username" binding:"required"`
	Password			string `json:"password" binding:"required"`
	PreferredTimeZone	string `json:"preferred_timezone" binding:"required,timezone"`
}

type LoginInput struct {
	Username		string `json:"username" binding:"required"`
	Password		string `json:"password" binding:"required"`
}
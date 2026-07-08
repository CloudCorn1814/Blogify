package user

type User struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"passwordhash"`
}

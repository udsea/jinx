package db

type User struct {
	ID           string
	DisplayName  string
	Email        string
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}


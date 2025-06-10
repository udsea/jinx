package db

import (
	"context"
	"log"
	"os"

	"github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *libsql.Client

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
	client, err := libsql.Open(dbURL, "")
	if err != nil {
		log.Fatalf("failed to connect to Turso: %v", err)
	}
	DB = client
}

func InsertUser(ctx context.Context, user User) error {
	_, err := DB.Exec(ctx, `
		INSERT OR REPLACE INTO users (id, display_name, email, access_token, refresh_token, expires_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		user.ID, user.DisplayName, user.Email, user.AccessToken, user.RefreshToken, user.ExpiresAt)
	return err
}


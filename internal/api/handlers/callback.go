package handlers

import (
	"context"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"

	"github.com/udbhav/jinx-backend/internal/db"
)

func SpotifyCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("No code in callback")
	}

	conf := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		Scopes:       []string{spotifyauth.ScopeUserReadEmail, spotifyauth.ScopePlaylistReadPrivate},
		Endpoint:     spotifyauth.Endpoint,
	}

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token exchange failed: " + err.Error())
	}

	httpClient := conf.Client(context.Background(), token)
	client := spotify.New(httpClient)

	user, err := client.CurrentUser(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user: " + err.Error())
	}

	// Insert into DB
	err = db.InsertUser(context.Background(), db.User{
		ID:           user.ID,
		DisplayName:  user.DisplayName,
		Email:        user.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.Expiry.Unix(),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to insert user: " + err.Error())
	}

	// Return JSON (or redirect)
	return c.JSON(fiber.Map{
		"user":         user.ID,
		"display_name": user.DisplayName,
		"email":        user.Email,
	})
}


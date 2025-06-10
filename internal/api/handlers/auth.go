package handlers

import (
	"github.com/udsea/jinx-backend/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func SpotifyLogin(c *fiber.Ctx) error {
	state := "jinx-state" // Ideally generate and verify CSRF-safe state
	authURL := auth.GetSpotifyAuthURL(state)
	return c.Redirect(authURL)
}


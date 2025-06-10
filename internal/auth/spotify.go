package auth

import (
	"fmt"
	"net/url"
	"os"
)

func GetSpotifyAuthURL(state string) string {
	params := url.Values{}
	params.Add("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	params.Add("response_type", "code")
	params.Add("redirect_uri", os.Getenv("SPOTIFY_REDIRECT_URI"))
	params.Add("scope", "playlist-read-private playlist-read-collaborative user-read-email")
	params.Add("state", state)

	return fmt.Sprintf("https://accounts.spotify.com/authorize?%s", params.Encode())
}


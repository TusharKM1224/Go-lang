package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"os" // for getting env variables

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     "<client ID>",
	ClientSecret: "<secret key>", //Get from env variables.
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func generateStateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateStateToken()
	url := googleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOnline)
	fmt.Println("OAuth2 URL:", url)                                                                                                     // Debugging
	http.SetCookie(w, &http.Cookie{Name: "oauthstate", Value: state, HttpOnly: true, Secure: false, SameSite: http.SameSiteStrictMode}) // store state
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")
	// state := r.URL.Query().Get("state")
	// cookie, err := r.Cookie("oauthstate")

	// if err != nil || cookie.Value != state {
	// 	http.Error(w, "Invalid state", http.StatusBadRequest)
	// 	return
	// }

	token, err := googleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to exchange token: %v", err), http.StatusInternalServerError)
		return
	}

	client := googleOAuthConfig.Client(ctx, token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user info: %v", err), http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse user info: %v", err), http.StatusInternalServerError)
		return
	}

	userJSON, err := json.MarshalIndent(userInfo, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal user info: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User Info: %s", userJSON)
}

func main() {
	http.HandleFunc("/auth/google/login", googleLoginHandler)
	http.HandleFunc("/auth/google/callback", googleCallbackHandler)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

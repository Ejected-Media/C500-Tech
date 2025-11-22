package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/discord"
)

// CONFIGURATION
// In production, fetch these from os.Getenv() for security
var conf = &oauth2.Config{
    ClientID:     "YOUR_CLIENT_ID_HERE",     // <-- PASTE FROM PORTAL
    ClientSecret: "YOUR_CLIENT_SECRET_HERE", // <-- PASTE FROM PORTAL
    RedirectURL:  "http://localhost:8080/auth/callback",
    Scopes:       []string{"identify", "email"},
    Endpoint:     discord.Endpoint,
}

func main() {
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/auth/callback", handleCallback)

    fmt.Println("C500 Auth Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// 1. THE LANDING PAGE
func handleHome(w http.ResponseWriter, r *http.Request) {
    html := `<html><body>
    <h1>Welcome to Classroom 500</h1>
    <p>Mentor Onboarding Portal</p>
    <a href="/login"><button>Login with Discord</button></a>
    </body></html>`
    w.Write([]byte(html))
}

// 2. THE REDIRECT (Send them to Discord)
func handleLogin(w http.ResponseWriter, r *http.Request) {
    // We generate the URL that asks Discord for permission
    url := conf.AuthCodeURL("state-token", oauth2.AccessTypeOnline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// 3. THE HANDSHAKE (They came back!)
func handleCallback(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    // A. Get the "code" Discord sent us in the URL query params
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Code not found", http.StatusBadRequest)
        return
    }

    // B. Exchange the "code" for an actual "Access Token"
    // This happens securely between our server and Discord's server
    token, err := conf.Exchange(ctx, code)
    if err != nil {
        http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // C. Use the Token to ask Discord "Who is this user?"
    client := conf.Client(ctx, token)
    resp, err := client.Get("https://discord.com/api/users/@me")
    if err != nil {
        http.Error(w, "Failed to get user info", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // D. Read and Display the User Data (For now)
    userData, _ := io.ReadAll(resp.Body)
    
    // In the next step, we will Parse this JSON and save it to a Cookie/Session
    w.Header().Set("Content-Type", "application/json")
    w.Write(userData)
}

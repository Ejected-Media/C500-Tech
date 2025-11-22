package main

import (
    "context"
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/discord"
)

// --- CONFIG ---
// Replace these with your actual keys from the Developer Portal!
var conf = &oauth2.Config{
    ClientID:     "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_CLIENT_SECRET", 
    RedirectURL:  "http://localhost:8080/auth/callback",
    Scopes:       []string{"identify", "email"},
    Endpoint:     discord.Endpoint,
}

// --- DATA MODELS ---
type DiscordUser struct {
    ID         string `json:"id"`
    Username   string `json:"username"`
    GlobalName string `json:"global_name"` 
    Avatar     string `json:"avatar"`
}

// --- TEMPLATE ---
// (Paste the 'const formHTML' string from section 2 here)
// For brevity, I am assuming the const is defined above.
var tmpl = template.Must(template.New("form").Parse(formHTML))

func main() {
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/auth/callback", handleCallback)
    http.HandleFunc("/submit-application", handleSubmit) // New Handler!

    fmt.Println("C500 Auth Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`<a href="/login">Login to Apply</a>`))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    url := conf.AuthCodeURL("random-state-string", oauth2.AccessTypeOnline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    // 1. Get Code
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "No code provided", http.StatusBadRequest)
        return
    }

    // 2. Exchange Code for Token
    token, err := conf.Exchange(ctx, code)
    if err != nil {
        http.Error(w, "Exchange failed", http.StatusInternalServerError)
        return
    }

    // 3. Fetch User Profile
    client := conf.Client(ctx, token)
    resp, err := client.Get("https://discord.com/api/users/@me")
    if err != nil {
        http.Error(w, "Fetch failed", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // 4. Parse JSON into Struct
    var user DiscordUser
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        http.Error(w, "JSON parse failed", http.StatusInternalServerError)
        return
    }

    // 5. Render HTML Form with User Data
    // This injects user.ID and user.Avatar into the HTML
    if err := tmpl.Execute(w, user); err != nil {
        http.Error(w, "Template render failed", http.StatusInternalServerError)
    }
}

// Placeholder for Step 3
func handleSubmit(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // Grab the form data
    philosophy := r.FormValue("philosophy")
    discordName := r.FormValue("discord_name")
    
    fmt.Fprintf(w, "Application Received for %s! Philosophy: %s", discordName, philosophy)
}

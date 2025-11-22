import "time"

type MentorApplication struct {
    DiscordID    string    `firestore:"discord_id"`
    DiscordName  string    `firestore:"discord_name"`
    TwitchURL    string    `firestore:"twitch_url"`
    Philosophy   string    `firestore:"philosophy"`
    WorkLink     string    `firestore:"work_link"`
    Status       string    `firestore:"status"` // e.g., "pending", "approved"
    SubmittedAt  time.Time `firestore:"submitted_at"`
}


import (
    "cloud.google.com/go/firestore"
    firebase "firebase.google.com/go/v4"
    "google.golang.org/api/option"
)


var fsClient *firestore.Client


func initFirestore() {
    ctx := context.Background()
    conf := &firebase.Config{ProjectID: "YOUR_PROJECT_ID"} // e.g. c500-tech

    // Note: On App Engine, you don't need option.WithCredentialsFile
    // For local dev, point to your serviceAccountKey.json
    opt := option.WithCredentialsFile("serviceAccountKey.json")
    
    app, err := firebase.NewApp(ctx, conf, opt)
    if err != nil {
        log.Fatalf("error initializing app: %v", err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalf("error getting firestore client: %v", err)
    }
    fsClient = client
}


import "bytes"

func sendDiscordNotification(app MentorApplication) {
    webhookURL := "YOUR_WEBHOOK_URL_HERE"

    // Format the message using Discord's JSON structure
    // We can mention roles using <@&ROLE_ID> if you want to ping Admins
    jsonBody := []byte(fmt.Sprintf(`{
        "content": "🚨 **New Mentor Application!**",
        "embeds": [{
            "title": "%s applied to be a mentor",
            "color": 5763719,
            "fields": [
                {"name": "Philosophy", "value": "%s"},
                {"name": "Twitch", "value": "%s"},
                {"name": "Link", "value": "%s"}
            ]
        }]
    }`, app.DiscordName, app.Philosophy, app.TwitchURL, app.WorkLink))

    http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    ctx := context.Background()

    // 1. Construct the Application Object
    app := MentorApplication{
        DiscordID:   r.FormValue("discord_id"),
        DiscordName: r.FormValue("discord_name"),
        TwitchURL:   r.FormValue("twitch_url"),
        Philosophy:  r.FormValue("philosophy"),
        WorkLink:    r.FormValue("work_link_1"),
        Status:      "pending",
        SubmittedAt: time.Now(),
    }

    // 2. Save to Firestore
    // We use .Add() to let Firestore generate a unique ID
    _, _, err := fsClient.Collection("mentors").Add(ctx, app)
    if err != nil {
        log.Printf("Failed to save application: %v", err)
        http.Error(w, "Database Error", http.StatusInternalServerError)
        return
    }

    // 3. Ping Discord (Fire and Forget)
    go sendDiscordNotification(app)

    // 4. Show Success Page
    w.Write([]byte(`
        <html>
        <body style="background:#2c2f33; color:white; text-align:center; padding-top:50px; font-family:sans-serif;">
            <h1>Application Received! 🚀</h1>
            <p>Thanks, ` + app.DiscordName + `. We have pinged the admins.</p>
            <p>You can close this window.</p>
        </body>
        </html>
    `))
}


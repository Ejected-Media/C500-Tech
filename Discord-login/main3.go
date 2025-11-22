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


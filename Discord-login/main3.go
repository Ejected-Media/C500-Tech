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


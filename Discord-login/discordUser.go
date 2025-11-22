// User represents the data we get back from Discord
type DiscordUser struct {
    ID            string `json:"id"`
    Username      string `json:"username"`
    Discriminator string `json:"discriminator"` // The #1234 part (legacy but still used)
    GlobalName    string `json:"global_name"`   // The new display name
    Avatar        string `json:"avatar"`
    Email         string `json:"email"`
}

This is the "Blueprint" for your entire ecosystem. It explains how the Android Game, the Web Portal, and the Discord Bot talk to each other without creating a mess of dependencies.
You can save this file as TECHNICAL_ARCHITECTURE.md in the root of your Ejected-Media organization. It is the document you would show to a new developer (or a future version of yourself) to explain "how it all works."
Technical Architecture: Ejected Media Ecosystem
Version: 1.0.0
Status: Active Development
Maintainer: Ejected Media
1. High-Level System Design
The system follows a Hub-and-Spoke model. The Firestore Database acts as the central hub (Source of Truth), while three distinct "Spokes" (Android, Web, Bot) read and write to it.
The Three Spokes
 * The Client (Drift Miner): A "Dumb" Android client. It plays the game and sends signed score packets. It never writes directly to the database; it talks to the Go API.
 * The Portal (C500.tech): A Go web server (App Engine) handling OAuth2 and User Onboarding. It verifies identity and writes "Applications" to Firestore.
 * The Enforcer (Discord Bot): A Go background process (GCE) that watches Firestore for changes (e.g., new applications) and manages Discord Channels/Roles.
2. Component Detail: Android Client (Drift Miner)
 * Architecture: MVVM (Model-View-ViewModel) + Clean Architecture.
 * Core Language: Kotlin.
 * Graphics: Custom SurfaceView with a manual Game Loop (60 FPS).
 * Physics: Deterministic "Inertial Drift" engine (No Box2D/Unity overhead).
 * Network: Retrofit (OkHttp) for REST API calls.
The Mobile Studio Workflow
Developed entirely on Android devices.
 * Code Entry: QuickEdit+ (Editor).
 * File Management: Files by Google (Structure verification).
 * Version Control: Spck Editor (Git Client).
 * Testing: Firefox Focus (Stateless Preview) + Chrome (Compatibility).
3. Component Detail: The Backend API (Go)
 * Platform: Google App Engine (Standard Environment).
 * Language: Go 1.22+.
 * Responsibility:
   * Validates Game Scores (Anti-Cheat logic).
   * Handles Discord OAuth2 Handshakes.
   * sanitizes inputs before writing to Firestore.
Sequence Flow: Mentor Intake
When a user applies to be a mentor on C500.tech:
 * User clicks "Login with Discord" on C500.tech.
 * Go Server redirects to Discord OAuth2 API.
 * Discord returns auth_code.
 * Go Server exchanges code for access_token + user_id.
 * Go Server saves Application Form to mentors collection in Firestore.
 * Go Server fires a Webhook to the Discord Admin Channel.
4. Component Detail: The Discord Bot
 * Platform: Google Compute Engine (e2-micro).
 * OS: Alpine Linux (Custom Setup).
 * Process Manager: OpenRC (No Systemd).
 * Binary: Static Go Binary (CGO_ENABLED=0).
Core Responsibilities
 * !officehours start: Creates a Stage Channel via Discord API, sets permissions so only the Mentor can speak, and announces the topic.
 * !review_latest: Fetches the newest document from the mentors Firestore collection and posts an interactive "Approve/Deny" card.
5. Data Architecture (Firestore Schema)
We use a NoSQL Document structure optimized for fast reads.
Collection: users
Public Profile Data
 * doc_id: discord_user_id
 * username: string
 * joined_at: timestamp
 * roles: array ["student", "mentor", "admin"]
Collection: scores
Drift Miner High Scores
 * user_id: string (Ref)
 * score: number (Integer)
 * level: number
 * platform: string ("android" | "ios")
 * timestamp: server_timestamp
Collection: mentors
Intake Applications
 * discord_id: string
 * philosophy: string (Text blob)
 * availability: map
   * timezone: string
   * can_do_live: boolean
 * status: enum ("pending", "approved", "rejected")
6. Deployment Pipelines
Android-to-Cloud (The "Termux Protocol")
How we ship server code from a phone.
 * Write Code: Edit main.go in QuickEdit+.
 * Compile: Open Termux.
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bot main.go

 * Ship:
   scp bot c500@<SERVER_IP>:~/

 * Restart:
   ssh c500@<SERVER_IP> "rc-service c500-bot restart"

YouTube Resource
This video perfectly explains the value of documenting your architecture like this, specifically using the "C4 Model" (Context, Containers, Components, Code) which we effectively used above.
Intro to Software Architecture
This video is relevant because it breaks down how to visualize and document software systems clearly, mirroring the structure we just built for Ejected Media.
YouTube video views will be stored in your YouTube History, and your data will be stored and used by YouTube according to its Terms of Service

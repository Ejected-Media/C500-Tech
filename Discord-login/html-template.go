import "html/template"

// A simple HTML form styled for C500
const formHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>C500 Mentor Application</title>
    <style>
        body { font-family: sans-serif; background: #2c2f33; color: white; display: flex; justify-content: center; padding: 50px; }
        .container { background: #23272a; padding: 40px; border-radius: 8px; width: 500px; box-shadow: 0 4px 15px rgba(0,0,0,0.5); }
        .profile { display: flex; align-items: center; margin-bottom: 20px; border-bottom: 1px solid #444; padding-bottom: 20px; }
        .avatar { width: 50px; height: 50px; border-radius: 50%; margin-right: 15px; }
        input, textarea { width: 100%; padding: 10px; margin-top: 5px; margin-bottom: 15px; background: #40444b; border: none; color: white; border-radius: 4px; }
        button { background: #7289da; color: white; padding: 12px 20px; border: none; border-radius: 4px; cursor: pointer; font-weight: bold; width: 100%; }
        button:hover { background: #5b6eae; }
        label { font-weight: bold; font-size: 0.9em; color: #b9bbbe; }
    </style>
</head>
<body>

<div class="container">
    <div class="profile">
        <img class="avatar" src="https://cdn.discordapp.com/avatars/{{.ID}}/{{.Avatar}}.png" alt="Avatar">
        <div>
            <h2>Welcome, {{.GlobalName}}</h2>
            <small>ID: {{.ID}}</small>
        </div>
    </div>

    <form action="/submit-application" method="POST">
        <input type="hidden" name="discord_id" value="{{.ID}}">
        <input type="hidden" name="discord_name" value="{{.GlobalName}}">

        <label>Twitch Channel URL</label>
        <input type="text" name="twitch_url" placeholder="https://twitch.tv/..." required>

        <label>Your CSS Philosophy (Why layout matters)</label>
        <textarea name="philosophy" rows="4" placeholder="I believe CSS Grid is..." required></textarea>

        <label>Link to your best CodePen / Work</label>
        <input type="text" name="work_link_1" placeholder="https://codepen.io/..." required>

        <button type="submit">Submit Application</button>
    </form>
</div>

</body>
</html>
`

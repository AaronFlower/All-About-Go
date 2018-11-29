<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width" />
        <title>Session count</title>
        <style type="text/css" media="screen">
            body {
                width: 60%;
                margin: auto;
            }
        </style>
    </head>
    <body>
        <h2>Hello GO!</h2>
        {{ if (.UserName) }}
        <p>Welcome <strong>{{.UserName}}</strong>, you have accessed our site {{.AccessTimes}} times!</p>
        <a href="/logout">Logout</a>
        {{ else }}
        <p>Welcome, please <a href="/login">Login</a></p>
        {{ end }}
    </body>
</html>

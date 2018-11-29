<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Form Submit</title>
    <style type="text/css" media="screen">
        body {
            width: 60%;
            margin: 20px auto;
        }
    </style>
</head>
<body>
    <h2>Please Login</h2>
    <form action="/login" method="post">
        User: <input type="text" name="username" value="{{.UserName}}"/>
        <br/>
        Password: <input type="password" name="password" />
        <br/>
        <input type="hidden" name="token" id="token" value="{{.Md5}}" />
        <input type="submit" value="Login" />
    </form>
</body>
</html>

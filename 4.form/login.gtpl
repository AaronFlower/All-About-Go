<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Form Submit</title>
</head>
<body>
    <form action="/login" method="post">
        <input type="checkbox" name="interest" value="football" />Football
        <input type="checkbox" name="interest" value="basketball" />Basketball
        <input type="checkbox" name="interest" value="Tennise" />Tennise
        User: <input type="text" name="username" />
        Password: <input type="password" name="password" />
        <input type="hidden" name="token" id="token" value="{{.}}" />
        <input type="submit" value="Login" />
    </form>
</body>
</html>

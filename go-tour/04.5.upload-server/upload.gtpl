<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>File Upload</title>
</head>
<body>
    <form action="/upload" method="post" enctype="multipart/form-data">
        <input type="file" name="uploadfile" />
        <input type="hidden" name="toke" value="{{.}}" />
        <input type="submit" value="upload" />
    </form>
</body>
</html>

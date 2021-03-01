package view

var LoginFormTmpl = string(`
<html>
	<body>
	<form action="/" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	<a href="/login">login</a>
	</body>
</html>
`)

var LogoutFormTmpl = string(`
<html>
	<body>
		<a href="/logout">logout</a>
	</body>
</html>
`)
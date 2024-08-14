package template

func GetForgetPasswordTemplate(token string) string {
	return `
	<h1>Reset Password</h1>
	<a href="http://localhost:3000/reset-password?key=` + token + `">Click here to reset your password</a>
	`
}

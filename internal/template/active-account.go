package template

func GetActiveAccountTemplate(token string) string {
	return `
	<h1>Verify your email</h1>
	<a href="http://localhost:3000/active-account?key=` + token + `">Click here to verify your email</a>
	`
}

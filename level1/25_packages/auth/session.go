package auth

func GetSession() string {
	return extractSession()
}

// private method / busines logic
func extractSession() string {
	return "logged in successfully"
}

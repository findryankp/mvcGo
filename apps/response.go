package apps

func CreateResponse() string {
	var text = `package helpers

func ResponseSuccess(message string, data any) map[string]any {
	return map[string]any{
		"status":  true,
		"message": message,
		"data":    data,
	}
}

func ResponseFail(message string) map[string]any {
	return map[string]any{
		"status":  false,
		"message": message,
	}
}
	`

	return text
}

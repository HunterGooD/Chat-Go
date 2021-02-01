package app

// ErrorRequest ошибки в случаи их возникновения
type ErrorRequest struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

// SignUpUser структура принимающая параметры для регистрации
type SignUpUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

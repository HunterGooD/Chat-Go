package app

// ErrorRequest ошибки в случаи их возникновения
type ErrorRequest struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

type userAuthResponse struct {
	Token string
}

// SignUpUser структура принимающая параметры для регистрации
type SignUpUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// SignInUser структура для входа в систему
type SignInUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

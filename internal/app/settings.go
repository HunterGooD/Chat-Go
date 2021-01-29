package app

// Config основные настройки
type Config struct {
	DB      DBConfig        `json:"db"`
	Secrets *PasswordConfig `json:"-"`
}

// DBConfig настройки для базы данных
type DBConfig struct {
	DBType     string `json:"db"` //  "mysql", "postres", "sqlite"
	DBName     string `json:"name"`
	DBUser     string `json:"user"`
	DBPassword string `json:"password"`
	DBHost     string `json:"host"`
	DBPort     string `json:"port"`
}

// PasswordConfig конфиг для паролей и ключей
type PasswordConfig struct {
	DBPassword string `json:"db_password"`
	AESKey     string `json:"aes_key"`
}

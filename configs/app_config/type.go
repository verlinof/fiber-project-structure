package app_config

type AppConfig struct {
	AppPort      string `env:"APP_PORT"`
	AppUrl       string `env:"APP_URL"`
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
}

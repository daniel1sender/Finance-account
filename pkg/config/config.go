package config

type Config struct {
	DatabaseURL string `envconfig:"DB_URL" required:"true"`
	Port        string `envconfig:"API_PORT" required:"true" default:":3000"`
	TokenSecret string `envconfig:"TOKEN_SECRET" required:"true"`
	ExpTime     string `envconfig:"EXP_TIME" required:"true"`
}

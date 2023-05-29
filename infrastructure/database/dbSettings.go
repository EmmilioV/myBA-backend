package database

type DBSettings struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     string `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DATABASE_NAME" required:"true"`
	SslMode  string `envconfig:"DATABASE_SSLMODE" required:"true"`
}

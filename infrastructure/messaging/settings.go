package messaging

type Settings struct {
	Url string `envconfig:"MESSAGING_URL" required:"true"`
}

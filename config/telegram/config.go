package telegram

import (
	"flag"
)

// Config defines telegram bot settings.
type Config struct {
	Token string
}

// Export returns telegram configuration.
func Export(flag *flag.FlagSet, t *Config) {
	flag.StringVar(&t.Token, "token", "", "Token to access the Teelgram HTTP API")
}

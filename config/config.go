package config

import (
	"flag"
	"log"
	"os"

	"github.com/gobwas/flagutil"
	"ia.samoylov/telegram.bot.sbmm/config/telegram"
)

// Config defines telegram bot settings.
type Config struct {
	Telegram telegram.Config
	Debug    bool
}

// Export returns telegram configuration.
func Export() *Config {
	flags := flag.NewFlagSet("telegram.bot.sbmm", flag.ExitOnError)

	var conf = Config{}

	flags.BoolVar(&conf.Debug, "debug", false, "Debug mode, the bot will send error messages to the user.")

	flagutil.Subset(flags, "telegram", func(sub *flag.FlagSet) {
		telegram.Export(sub, &conf.Telegram)
	})

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Panic(err)
		os.Exit(0)
	}

	return &conf
}

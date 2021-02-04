package main

import (
	"fmt"

	"ia.samoylov/telegram.bot.sbmm/config"
)

func main() {
	var config = *config.Export()

	fmt.Println(config.Telegram.GoString())
	fmt.Println(config.GoString())
}

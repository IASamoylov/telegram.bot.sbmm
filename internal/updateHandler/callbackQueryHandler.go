package updatehandler

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CallbackQueryHandle message from telegram bot.
func (ctx *HandlerContext) CallbackQueryHandle(update *tgbotapi.Update) {
	data := strings.Split(update.CallbackQuery.Data, ":")
	if len(data) != 0 {
		switch data[0] {
		case "platform":
			ctx.handlePlatformCallbackQuery(update, data)
		default:
			log.Printf(`Unknown command: %v`, data[0])
		}
	}
}

func (ctx *HandlerContext) handlePlatformCallbackQuery(update *tgbotapi.Update, data []string) {
	platform := data[1]
	user, err := ctx.UserStorage.Get(update.CallbackQuery.From.ID)

	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

	value, platformIsExist := platforms[platform]
	fmt.Println(value)

	if !platformIsExist {
		msg.Text = fmt.Sprintf(`Выбраная платформа %s не существует. `, platform)
	}

	user.Platform = platform
	ctx.UserStorage.Update(user)

	msg.Text = fmt.Sprintf(`Выбраная платформа %s Была сохранена. Посмотреть текущую выбранную платформу можно с помощью команды /settings`, value)

	ctx.Bot.Send(msg)
}

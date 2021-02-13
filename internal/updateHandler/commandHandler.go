package updatehandler

import (
	"bytes"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"ia.samoylov/telegram.bot.sbmm/config"
	"ia.samoylov/telegram.bot.sbmm/internal/models"
)

// HandlerContext define command processing context.
type HandlerContext struct {
	Bot         *tgbotapi.BotAPI
	UserStorage models.UserStorage
	AppConfig   *config.Config
}

// CommandHandle message from telegram bot.
func (ctx *HandlerContext) CommandHandle(update *tgbotapi.Update) {
	if ctx.AppConfig.Debug {
		ctx.handlerDebugCode(update)
	}

	switch update.Message.Command() {
	case "start":
		ctx.handleStartCommand(update)
	case "help":
		ctx.handleHelpCommand(update)
	case "settings":
		ctx.handleSettingsCommand(update)
	case "reset":
		ctx.handleResetCommand(update)
	case "platform":
		ctx.handlePlatformCommand(update)
	case "username":
		ctx.handleResetCommand(update)
	default:
		ctx.handleDefaultCommand(update)
	}
}

func (ctx *HandlerContext) handleStartCommand(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	user, err := ctx.UserStorage.Get(update.Message.From.ID)

	if err != nil {
		log.Println(err)
		err := ctx.UserStorage.Add(models.User{
			ID:       update.Message.From.ID,
			Language: update.Message.From.LanguageCode,
		})

		if err != nil {
			log.Println(err)
			msg.Text = "Мне очень жаль, но по какой-то из причин я не могу сейчас обработать твою команду, попробуй еще раз чуть позднее."
		}
	}

	var msgBytes bytes.Buffer

	msgBytes.WriteString(fmt.Sprintf(`Приветствую мой друг, %s! Меня зовут @%s. Я создан для того чтобы быстро и удобно получать статистику с сайта https://sbmmwarzone.com.`, update.Message.From.FirstName, ctx.Bot.Self.UserName))
	msgBytes.WriteString("\n\n")

	if len(user.Platform) == 0 && len(user.UserName) == 0 {
		msgBytes.WriteString("Перед тем как мы приступим к работе мне нужно, чтобы ты мне немного помог.")

		msgBytes.WriteString(" ")
		if len(user.Platform) == 0 {
			msgBytes.WriteString("Необходимо указать платформу на которой играешь, выбрать платформу можно c помощью комнады /platform.")
		}

		msgBytes.WriteString(" ")
		if len(user.UserName) == 0 {
			msgBytes.WriteString("Необходимо указать никнем c помощью комнады /username или сразу указав никнейм введя команду и следом имя /username {YOU_NICK_NAME}")
		}
	}

	msgBytes.WriteString("\n\n")
	msgBytes.WriteString("Если тебе потребуется помощь, то ты можешь воспользовать командой /help, где я тебе подробно расскажу о всех своих возможностях.")

	msg.Text = msgBytes.String()

	ctx.Bot.Send(msg)
}

func (ctx *HandlerContext) handleResetCommand(update *tgbotapi.Update) {
	ctx.UserStorage.Delete(update.Message.From.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Настройки пользователя сброшены")

	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

	ctx.Bot.Send(msg)
}

var platforms = map[string]string{
	"psn":    "PlayStation",
	"xbl":    "Xbox Live",
	"battle": "Battle.net",
}

func (ctx *HandlerContext) handlePlatformCommand(update *tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for key, value := range platforms {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(value, fmt.Sprintf(`platform:%s`, key))
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}
	msg.ReplyMarkup = keyboard
	msg.Text += "Выбирите платформу на которой играете. Платформу можно изменить в любой момент снова написав данную команду."

	ctx.Bot.Send(msg)
}

func (ctx *HandlerContext) handleHelpCommand(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Команда /help находится в разработке.")
	ctx.Bot.Send(msg)
}

func (ctx *HandlerContext) handleSettingsCommand(update *tgbotapi.Update) {
	user, err := ctx.UserStorage.Get(update.Message.From.ID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if err != nil {
		log.Println(err)
		msg.Text = "Настройки текущего пользователя не найдены, возможно вы ниразу не выполняли команду /start либо сбросили свои настройки с помощью команды /reset."
	} else {
		msg.Text = fmt.Sprintf(`Выбраная платформа: %s, Выбраный пользователь: %s, Выбраный язык: %s`, platforms[user.Platform], user.UserName, user.Language)

	}

	ctx.Bot.Send(msg)
}

func (ctx *HandlerContext) handleDefaultCommand(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Данная команда не поддерживается, воспользуйтесь командой /help чтобы узнать актуальный список команд.")
	ctx.Bot.Send(msg)
}

func (ctx *HandlerContext) handlerDebugCode(update *tgbotapi.Update) {
	message := fmt.Sprintf(`ChatID: %d, User: %d, Command: /%s, Arguments: %s`, update.Message.Chat.ID, update.Message.From.ID, update.Message.Command(), update.Message.CommandArguments())
	log.Println(message)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	ctx.Bot.Send(msg)
}

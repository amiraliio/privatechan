package controllers

import (
	"github.com/amiraliio/tgbp/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

func onActionEvents(app *config.App, bot *tb.Bot) {

	bot.Handle(tb.OnChannelPost, func(message *tb.Message) {
		if generalEventsHandler(app, bot, message, &Event{
			Event:      tb.OnChannelPost,
			UserState:  config.LangConfig.GetString("STATE.REGISTER_CHANNEL"),
			Command:    config.LangConfig.GetString("COMMANDS.ENABLE_CHAT"),
			Controller: "RegisterChannel",
		}) {
			Init(app, bot, true)
		}
	})

	bot.Handle(tb.OnAddedToGroup, func(message *tb.Message) {
		if generalEventsHandler(app, bot, message, &Event{
			Event:      tb.OnAddedToGroup,
			UserState:  config.LangConfig.GetString("STATE.REGISTER_GROUP"),
			Controller: "RegisterGroup",
		}) {
			Init(app, bot, true)
		}
	})

	bot.Handle(tb.OnNewGroupTitle, func(message *tb.Message) {
		if generalEventsHandler(app, bot, message, &Event{
			Event:      tb.OnNewGroupTitle,
			UserState:  config.LangConfig.GetString("STATE.UPDATE_GROUP_TITLE"),
			Controller: "UpdateGroupTitle",
		}) {
			Init(app, bot, true)
		}
	})

	bot.Handle(&addAnonMessage, func(message *tb.Message) {
		if generalEventsHandler(app, bot, message, &Event{
			Event:      &addAnonMessage,
			UserState:  config.LangConfig.GetString("STATE.ADD_ANON_MESSAGE"),
			Controller: "AddAnonMessageToChannel",
		}) {
			Init(app, bot, true)
		}
	})

}

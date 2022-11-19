package travel

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *DemoSubdomainCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"You can control me by sending these commands:\n\n"+
			"/new - create a new entity\n"+
			"/get - get an entity\n"+
			"/list - get a list of your entities\n"+
			"/edit - edit your entity\n"+
			"/delete - delete an existing entity\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}

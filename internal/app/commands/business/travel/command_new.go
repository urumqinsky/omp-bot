package travel

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *DemoSubdomainCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	c.subdomainService.Add(args)
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%s was added\n", args),
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}

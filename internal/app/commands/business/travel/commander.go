package travel

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type DemoSubdomainCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *subdomain.Service
}

func NewDemoSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *DemoSubdomainCommander {
	subdomainService := subdomain.NewService()

	return &DemoSubdomainCommander{
		bot:              bot,
		subdomainService: subdomainService,
	}
}

func (c *DemoSubdomainCommander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf(
			"DemoSubdomainCommander.HandleCallback: unknown callback name: %s",
			callbackPath.CallbackName,
		)
	}
}

func (c *DemoSubdomainCommander) HandleCommand(
	msg *tgbotapi.Message,
	commandPath path.CommandPath,
) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "new":
		c.New(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}

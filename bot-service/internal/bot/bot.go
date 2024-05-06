package bot

import (
	"context"
	tele "gopkg.in/telebot.v3"
	"log"
	"main/internal/config"
)

type Bot struct {
	Settings tele.Settings
	Instance *tele.Bot
	Ctx      context.Context
}

func New(ctx context.Context, config config.BotConfig) (*Bot, error) {
	settings := tele.Settings{
		Token: config.GetToken(),
		Poller: &tele.LongPoller{
			Timeout: config.GetPollerTimeout(),
		},
	}

	b, err := tele.NewBot(settings)
	if err != nil {
		return nil, err
	}

	b.Handle("/start", Start)
	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("help")
	})

	return &Bot{settings, b, ctx}, nil
}

func (b *Bot) Start() {
	log.Println("Starting TG Bot instance")
	b.Instance.Start()
}

func (b *Bot) Shutdown() {
	log.Println("TG Bot instance stopped")
	b.Instance.Stop()
}

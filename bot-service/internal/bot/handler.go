package bot

import tele "gopkg.in/telebot.v3"

func Start(c tele.Context) error {

	return c.Send(c.Sender().FirstName)
}

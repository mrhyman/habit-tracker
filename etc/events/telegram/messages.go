package telegram

const msgHelp = `I can save and keep your pages of interest. Also I can send them back on your demand.

In order to save page just send me a link

In order to get a random saved page send me /rnd command
Caution! This page will be removed after it!`

const msgHello = "Hi there! 🥸 \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command! 🤦‍♂️"
	msgNoSavedPages   = "Sorry, there is no saved page yet 🔎"
	msgSaved          = "Saved! ✅"
	msgAlreadyExists  = "This page is already in your list 🙈"
)

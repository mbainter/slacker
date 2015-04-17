package inithooks

import (
	sl "github.com/mbainter/slacker/slackerlib"
)

var Bai = sl.ShutdownHook{
	Name: `Bai`,
	Usage:`log a friendly goodbye message on a graceful shutdown`,
	Run:	func(bot *sl.Sbot){
		bot.Say(`I have been remotely shut down. Goodbye.`,`baymax-test`)
		sl.Logger.Info(`Caught SigTerm, slacker shutting down...ZOMGBAI!!`)
	},
}

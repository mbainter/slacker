package main

import(
	sl "github.com/mbainter/slacker/slackerlib"
	"github.com/mbainter/slacker/inithooks"
	"github.com/mbainter/slacker/handlers"
	"github.com/mbainter/slacker/chores"
)

func initPlugins(b *sl.Sbot) error{
	//init hooks
	b.Register(inithooks.Hai)
	b.Register(inithooks.Bai)
	//b.Register(inithooks.CbTest)
	//handlers
	b.Register(handlers.Syn)
	b.Register(handlers.Help)
	b.Register(handlers.Braintest)
	b.Register(handlers.ListChores)
	b.Register(handlers.ManageChores)
	b.Register(handlers.ChannelID)
	b.Register(handlers.Gifme)
	b.Register(handlers.Tableflip)
	b.Register(handlers.LoveAndWar)
	b.Register(handlers.IKR)
	b.Register(handlers.MetaUpdater)
	//b.Register(handlers.MetaInfo)
	//chores
	b.Register(chores.RTMPing)
	return nil
}

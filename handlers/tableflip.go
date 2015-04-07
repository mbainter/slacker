package handlers

import (
    "time"
    "math/rand"
    sl "github.com/mbainter/slacker/slackerlib"
)

var Tableflip = sl.MessageHandler{
	Name: `Tableflip`,
	Usage: `bot flips a unicode table whenever it overhears '(table)*flip(table)*'`,
	Method:  `HEAR`,
	Pattern: `(?i)(table)*flip(table)*`,
	Run: func(e *sl.Event, match []string) {
		flips := []string{
            "(╯°□°）╯︵ ┻━┻`",
            "(╯°□°）╯︵ ┻━┻'",
            "(ノ ゜Д゜)ノ ︵ ┻━┻'",
            "(╯°□°)╯︵ ┻━┻ ︵ ╯(°□° ╯)",
            "(ノ^_^)ノ┻━┻ ┬─┬ ノ( ^_^ノ)",
            "(╯°Д°）╯︵ /(.□ . \\)",
            "(╯'□')╯︵ ┻━┻",
            "(ﾉಥДಥ)ﾉ︵┻━┻･/",
        }
	 	now:=time.Now()
	    rand.Seed(int64(now.Unix()))
	    e.Respond(flips[rand.Intn(len(flips)-1)])
	},
}

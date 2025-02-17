package check

import (
	"log"
	"strings"

	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func Rate(data string) *o.Rate {
	sl_rate := strings.Split(data, " ")
	if len(sl_rate) != 4 {
		log.Println("check.Rate len != 4")
		return nil
	}
	var rate = make([]float64, 0, 4)
	for in := 0; in < 4; in++ {
		if sl.CheckNumber(sl_rate[in]) {
			rate = append(rate, sl.ParF(sl_rate[in]))
		} else {
			log.Println("check.Rate msg[in] == is not a number")
			return nil
		}
	}

	ev := o.NewEv()
	ev.SetEv(rate[0], rate[1], rate[2], 0)
	ev.SetPortion(rate[3])
	ev.SetWeight(rate[3])

	return o.NewRate(*ev)
}

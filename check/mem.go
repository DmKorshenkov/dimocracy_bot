package check

import (
	"fmt"
	"time"

	"github.com/DmKorshenkov/helper/bot/ymd"
)

func CheckMem(data string) int {
	var date int
	var ok bool
	//
	if date, ok = func() (int, bool) {
		t, err := time.Parse("01/02/06", data)
		if err != nil {
			//log.Println(err.Error())
			return 0, false
		}

		date = ymd.ConvDate_ymd(t.Year()%100, int(t.Month()), t.Day())
		return date, true
	}(); ok {
		fmt.Println("date 02/01/06 == ", date)
		return date
	}
	//
	if date, ok = func() (int, bool) {
		t, err := time.Parse("02/06", data)
		if err != nil {
			//log.Println(err.Error())
			return 0, false
		}

		date = ymd.ConvDate_ymd(t.Year()%100, int(t.Month()), 0)
		return date, true
	}(); ok {
		fmt.Println("date 01/06 == ", date)
		return date
	}
	if date, ok = func() (int, bool) {
		t, err := time.Parse("06", data)
		if err != nil {
			//log.Println(err.Error())
			return 0, false
		}

		date = ymd.ConvDate_ymd(t.Year()%100, 0, 0)
		return date, true
	}(); ok {
		fmt.Println("date 01/06 == ", date)
		return date
	}

	return 0
}

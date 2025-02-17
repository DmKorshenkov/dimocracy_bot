package ymd

import (
	"time"
)

func ConvDateYMD(i int) (y int, m int, d int) {
	y = i / 10000
	m = (i - y*10000) / 100
	d = (i - y*10000 - m*100)
	if y == 10 {
		y = 0
	}
	return y, m, d
}

func ConvDate_ymd(y, m, d int) (ymd int) {
	if y < 0 || m < 0 || d < 0 {
		return
	}
	return y*10000 + m*100 + d
}

func ConvDateNow() int {
	t := time.Now()
	data := (t.Year() % 100 * 10000) + (int(t.Month()) * 100) + (t.Day())

	return data
}

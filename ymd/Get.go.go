package ymd

import (
	"fmt"
	"log"
)

func Get[K comparable, V any](mp map[int]map[int]map[int]map[K][]V, date int) (any, string) {
	y, m, d := ConvDateYMD(date)
	nowY, nowM, nowD := ConvDateYMD(ConvDateNow())
	fmt.Println()

	if y != 0 {
		mpY, ok := getY(mp, y)
		if !ok {
			return nil, "Упс, в календаре нет такого года"
		}
		if m == 0 {
			return mpY, ""
		}
		mpM, ok := getM(mpY, m)
		if !ok {
			return nil, "Упс, в календаре нет такого месяца"
		}
		if d == 0 {
			return mpM, ""
		}
		mpD, ok := getD(mpM, d)
		if !ok {
			return nil, "Упс, в календаре нет такого day"
		}
		return mpD, ""
	} else {

		if m == 0 {
			mpY, ok := getY(mp, nowY)
			if !ok {
				return nil, "Упс, в календаре нет такого year"
			}
			mpM, ok := getM(mpY, nowM)
			if !ok {
				return nil, "Упс, в календаре нет такого month"
			}
			mpD, ok := getD(mpM, nowD)
			if !ok {
				return nil, "Упс, в календаре нет такого year"
			}
			return mpD, ""
		}
		mpM, _ := getM(mp[nowY], m)
		return mpM, ""
	}

}

func getD[K comparable, V any](mp map[int]map[K][]V, day int) (map[K][]V, bool) {
	if day, ok := mp[day]; ok {
		return day, ok
	} else {
		log.Println("mp[day] == false(not found in map)")
		return nil, ok
	}
}

func getM[K comparable, V any](mp map[int]map[int]map[K][]V, m int) (map[int]map[K][]V, bool) {
	if month, ok := mp[m]; ok {
		return month, ok
	} else {
		log.Println("mp[month]==false(not found in map)")
		return nil, ok
	}
}

func getY[K comparable, V any](mp2 map[int]map[int]map[int]map[K][]V, y int) (map[int]map[int]map[K][]V, bool) {
	if year, ok := mp2[y]; ok {
		return year, ok
	} else {
		log.Println("mp[year]==false(not found in map)")
		return nil, ok
	}
}

func getDay[K comparable, V any](mp2 map[int]map[int]map[int]map[K][]V, day int) map[K][]V {
	y, m, _ := ConvDateYMD(ConvDateNow())
	if mp2[y] == nil || mp2[y][m] == nil || mp2[y][m][day] == nil {
		fmt.Println("mp[day]==false(not found in map)")
		return nil
	}
	return mp2[y][m][day]
}

func getMonth[K comparable, V any](mp2 map[int]map[int]map[int]map[K][]V, m int) map[int]map[K][]V {
	y, _, _ := ConvDateYMD(ConvDateNow())
	if mp2[y] == nil || mp2[y][m] == nil {
		fmt.Println("mp[month]==false(not found in map)")
		return nil
	}
	return mp2[y][m]
}

func getYear[K comparable, V any](mp2 map[int]map[int]map[int]map[K][]V, y int) map[int]map[int]map[K][]V {
	if mp2[y] == nil {
		fmt.Println("mp[year]==false(not found in map)")
		return nil
	}
	return mp2[y]
}

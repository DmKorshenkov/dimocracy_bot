package in

import (
	"fmt"
	"strings"

	"github.com/DmKorshenkov/helper/bot/check"
	"github.com/DmKorshenkov/helper/bot/o"
)

type I struct {
	cmd  string
	key  string
	req  int32
	data string
}

func NewI() *I {
	return &I{}
}

func (i *I) NewI(cmd string, key string) *I {
	//i.cmd = strings.TrimSpace(cmd)
	i.cmd = strings.ToLower(strings.TrimSpace(cmd))
	i.key = helpTrim(key)
	//i.data = data
	//i.key = strings.TrimSpace(key)
	return i
}

func (i *I) CheckCmd() {
	switch i.cmd {
	case "rem":
		i.req = 'r'
		return
	case "mem":
		i.req = 'm'
		return
	case "cal":
		i.req = 'c'
		return
	case "запомни":
		i.req = 'r'
		return
	case "вспомни":
		i.req = 'm'
		return
	case "калькулятор":
		i.req = 'c'
		return
	default:
		return
	}
}

func (i *I) CheckKey() {
	switch i.key {
	case "weight":
		i.req += 'w'
		return
	case "prod":
		i.req += 'f'
		return
	case "rate":
		i.req += 'R'
		return
	case "meal take":
		i.req += 'm'
		return
	case "вес":
		i.req += 'w'
		return
	case "продукт":
		i.req += 'f'
		return
	case "норму бжу":
		i.req += 'R'
		return
	case "прием пищи":
		i.req += 'm'
		return
	default:
		i.req = 0
		return
	}
}

func (i *I) Check() {
	i.CheckCmd()
	i.CheckKey()
}

func (i *I) PI() {
	fmt.Println("input:")
	fmt.Println("i.cmd = ", i.cmd)
	fmt.Println("i.key = ", i.key)
}

func In(ch chan string, msg string) {

	var i = NewI()

	checK, got := func() ([]string, []string) {
		get := strings.SplitN(msg, "\n", 2)
		cmdkey := strings.SplitN(get[0], " ", 2)
		return cmdkey, get
	}()
	if len(checK) != 2 {
		ch <- "Упс)\nНеверное количество ключевых слов в команде.\n"
		close(ch)
		return
	}
	i.NewI(checK[0], checK[1])
	i.PI()
	i.Check()
	if i.req == 0 {
		ch <- "cmd/key!=true\n"
		close(ch)
		return
	} else {
		ch <- ("cmd/key==true")
	}

	if len(got) > 1 {
		i.data = strings.ToLower(got[1])
	}

	switch i.req {
	case 'r' + 'w':
		//
		weight := check.RemWeight(i.data)
		if weight == nil {
			ch <- "(check.RemWeight==nil)\nОдин раз - не пидорас.\nНо, Братан, ты пидр!\nДавай! Попробуй еще раз!"
		} else {
			ch <- o.RemWeight(*weight)
			close(ch)
			return
		}
	case 'm' + 'w':
		// mem weight
	case 'r' + 'f':
		/*/
		food := check.RemFood(i.data)
		if len(food) == 0 {
			return "(check.RemFood==len_0)\nОдин раз - не пидорас.\nНо, Братан, ты пидр!\nДавай! Попробуй еще раз!"
		} else {
			o.RemFood(food...)
		}/*/
	case 'r' + 'R':
		/*/
		rate := check.Rate(i.data)
		if rate == nil {
			return "(check.Rate==nil)\nОдин раз - не пидорас.\nНо, Братан, ты пидр!\nДавай! Попробуй еще раз!"
		} else {
			o.RemRate(*rate)
		}/*/
	case 'm' + 'R':
		//mem rate

	case 'r' + 'm':
		/*/
		prod := check.Prod(i.data)
		if len(prod) == 0 {
			return "(check.Prod==len_0)\nОдин раз - не пидорас.\nНо, Братан, ты пидр!\nДавай! Попробуй еще раз!"
		} else {
			meal := fnc.MealTake(prod...)
			fnc.RemMeal(meal)
		}/*/
	case 'm' + 'm':
		//mem meal

	}

	return
}

func helpTrim(str string) string {
	return strings.ToLower(strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(str, " "), "\n"), " "))
}

package check

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/sl"
)

func Prod(data string) []o.Prod {
	var sl2 = strings.Split(data, "\n")
	var prods = make([]o.Prod, 0, len(sl2))

	for _, str := range sl2 {
		var slStr = strings.Split(str, " ")
		if len(slStr) != 2 {
			log.Println("CheckProd len == false")
			return nil
		}
		if !sl.CheckNumber(slStr[1]) {
			log.Println("CheckProd Number == false")
			return nil
		}
		dir, _ := os.Getwd()
		fmt.Println()
		fmt.Println("in check prod -- ", dir)
		if o.MemFood(slStr[0]) == nil {
			log.Println("CheckProd MemFood == msg[i] not found in list foods")
			return nil
		}
		weight, _ := strconv.ParseFloat(slStr[1], 64)
		prods = append(prods, o.NewProd().SetProd(slStr[0], weight))
	}

	return prods
}

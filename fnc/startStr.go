package fnc

import (
	"log"
	"os"
)

func StartStr() string {
	err := os.Mkdir("DataBase", 0755)
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	} else {
		log.Println("success")
	}
	err = os.Chdir("./DataBase")
	if err != nil {
		log.Println(err.Error())
		return err.Error()
	} else {
		log.Println("success")
	}
	return "start is success"
}

package fnc

import (
	"log"
	"os"
)

func Start() string {
	os.Mkdir("DataBase", 0755)
	err := os.Chdir("./DataBase")
	dir, _ := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return err.Error() + "\n" + dir
	}
	return dir
}

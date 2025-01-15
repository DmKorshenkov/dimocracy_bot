package fnc

import (
	"log"
	"os"
)

func Start() string {
	err := os.Chdir("./Data")
	if err != nil {
		dir, _ := os.Getwd()
		log.Println(err.Error(), "\n", dir)
	}
	os.Mkdir("DataBase", 0755)
	err = os.Chdir("./DataBase")
	dir, _ := os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return err.Error() + "\n" + dir
	}
	return dir
}

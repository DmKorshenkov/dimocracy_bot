package fnc

import (
	"log"
	"os"
)

func Start() string {
	dirs, err := os.ReadDir("./")
	var str string
	if err != nil {
		wd, _ := os.Getwd()
		log.Println(err.Error(), "\n", wd)
	}
	for _, dir := range dirs {
		str += dir.Name() + "\n"
		log.Println(dir.Name())
		if dir.Name() == "Data" || dir.Name() == "data" {
			os.Chdir(dir.Name())
			wd, _ := os.Getwd()
			log.Println(wd)
			err := os.Mkdir("DataBase", 0755)
			if err != nil {
				return err.Error() + "\n" + wd
			}
			err = os.Chdir("./DataBase")
			wd, _ = os.Getwd()
			if err != nil {
				return err.Error() + "\n" + wd
			}
		}
	}
	wd, _ := os.Getwd()
	return wd + "\n" + str
}

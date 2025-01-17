package fnc

import (
	"log"
	"os"
)

func Start(ch chan string) {
	err := os.Mkdir("DataBase", 0755)
	if err != nil {
		log.Println(err.Error())
		ch <- err.Error()
	} else {
		ch <- "success"
	}
	err = os.Chdir("./DataBase")
	if err != nil {
		log.Println(err.Error())
		ch <- err.Error()
	} else {
		ch <- "success"
	}
	_, err = os.Create("rate.json")
	if err != nil {
		log.Println(err.Error())
		ch <- err.Error()
	} else {
		ch <- "success"
	}
	_, err = os.Create("rate_tmp.json")
	if err != nil {
		log.Println(err.Error())
		ch <- err.Error()
	} else {
		ch <- "success"
	}
	_, err = os.Create("weight.json")
	if err != nil {
		log.Println(err.Error())
		ch <- err.Error()
	} else {
		ch <- "success"
	}
	_, err = os.Create("product.json")
	if err != nil {
		log.Println(err.Error())
		ch <- err.Error()
	} else {
		ch <- "success"
	}
	_, err = os.Create("meal_take.json")
	if err != nil {
		log.Println(err.Error())
		ch <- err.Error()
	} else {
		ch <- "success"
	}
	close(ch)
}

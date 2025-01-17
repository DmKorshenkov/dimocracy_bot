package o

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DmKorshenkov/helper/bot/ymd"
)

type Weight struct {
	Info   string  `json:",omitempty"`
	Weight float64 `json:",omitempty"`
}

func NewW() *Weight {
	return &Weight{}
}
func (w *Weight) Look() string {
	var str string
	if w.Weight != 0 {
		str += fmt.Sprintf("Вес - %-4.2f", w.Weight)
	}
	if w.Info != "" {
		str += fmt.Sprintf("%s", w.Info)
	}
	return str
}

func NewSetW(weight float64, info string) *Weight {
	return &Weight{Weight: weight, Info: info}
}

func (w *Weight) SetWeight(weight float64) *Weight {
	w.Weight = weight
	return w
}

func (w *Weight) SetInfo(info string) {
	w.Info = info
}

func (w *Weight) Str() string {
	if w.Info == "" {
		return fmt.Sprintf("weight - %v", w.Weight)
	}
	return fmt.Sprintf("weight - %v\ninfo - %v", w.Weight, w.Info)
}

func RemWeight(w Weight) {
	f, err := os.OpenFile("weight.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}
	mp := make(map[int]map[int]map[int]map[string][]Weight)
	data, _ := os.ReadFile("weight.json")
	if len(data) != 0 {
		_ = json.Unmarshal(data, &mp)
		fmt.Println(mp)
	}
	_ = ymd.Set[string, Weight](mp, w.Info, w)
	fmt.Println(mp)
	data, err = json.MarshalIndent(mp, "", "	")
	if err != nil {
		fmt.Println("!")
	}
	n, _ := f.Write(data)
	if n != len(data) {
		fmt.Println("n!=data")
	}
	_ = f.Close()
}

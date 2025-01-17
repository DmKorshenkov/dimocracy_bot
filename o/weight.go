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

func RemWeight(w Weight) string {
	f, err := os.OpenFile("weight.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err.Error()
	}
	mp := make(map[int]map[int]map[int]map[string][]Weight)
	data, err := os.ReadFile("weight.json")
	if err != nil {
		return err.Error()
	}
	if len(data) != 0 {
		err = json.Unmarshal(data, &mp)
		if err != nil {
			return err.Error()
		}
	}
	_ = ymd.Set[string, Weight](mp, w.Info, w)
	data, err = json.MarshalIndent(mp, "", "	")
	if err != nil {
		return err.Error()
	}
	n, err := f.Write(data)
	if n != len(data) || err != nil {
		return "n, err != f.Write(data) RemWeight 72"
	}
	err = f.Close()
	if err != nil {
		return err.Error()
	}
	return "rem weight success\n"
}

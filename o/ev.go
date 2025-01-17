package o

import (
	"fmt"
	"math"
)

type Ev struct {
	P   float64 `json:"Prot"`
	F   float64 `json:"Fat"`
	C   float64 `json:"Carb"`
	Fb  float64 `json:"Fiber,omitempty"`
	Cal float64 `json:"Cal"`
	W   Weight  `json:"W,omitempty"`
}

func NewEv() *Ev {
	return &Ev{}
}

func (ev *Ev) Look() string {
	var str string
	ev.Round()
	if ev.W.Weight != 0 {
		str += fmt.Sprintf("Вес       - %-4.2f gram\n", ev.W.Weight)
	}
	if ev.P != 0 {
		str += fmt.Sprintf("Белков    - %-4.2f gram\n", ev.P)
	}
	if ev.F != 0 {
		str += fmt.Sprintf("Жиров     - %-4.2f gram\n", ev.F)
	}
	if ev.C != 0 {
		str += fmt.Sprintf("Углеводов - %-4.2f gram\n", ev.C)
	}
	if ev.Fb != 0 {
		str += fmt.Sprintf("Клетчатки - %-4.2f gram", ev.Fb)
	}
	if ev.Cal != 0 {
		str += fmt.Sprintf("Калорий   - %-4.2f\n", ev.Cal)
	}

	return str
}
func SetEv(p float64, f float64, c float64, fb float64) *Ev {
	return &Ev{P: p, F: f, C: c, Fb: fb, Cal: p*4 + f*9 + c*4 + fb*1.2, W: *NewSetW(100, "")}
}
func (ev *Ev) SetEv(p float64, f float64, c float64, fb float64) *Ev {
	ev.P = p
	ev.F = f
	ev.C = c
	ev.Fb = fb
	ev.Cal = p*4 + f*9 + c*4 + fb*1.2
	ev.W.SetWeight(100)
	return ev
}

func (ev *Ev) SetWeight(weight float64) *Ev {
	ev.W.Weight = weight
	return ev
}
func (Ev *Ev) SetOneGram() *Ev {
	Ev.W.Weight = 100
	Ev.P = ((Ev.P) / Ev.W.Weight)
	Ev.F = ((Ev.F) / Ev.W.Weight)
	Ev.C = ((Ev.C) / Ev.W.Weight)
	Ev.Fb = ((Ev.Fb) / Ev.W.Weight)
	Ev.Cal = ((Ev.Cal) / Ev.W.Weight)
	Ev.W.Weight = 1
	//Ev.Round()
	return Ev
}
func (Ev *Ev) SetPortion(weight float64) *Ev {

	Ev.W.Weight = weight
	Ev.P = (Ev.P) * Ev.W.Weight
	Ev.F = (Ev.F) * Ev.W.Weight
	Ev.C = (Ev.C) * Ev.W.Weight
	Ev.Fb = (Ev.Fb) * Ev.W.Weight
	Ev.Cal = (Ev.Cal) * Ev.W.Weight
	Ev.Round()
	return Ev
}
func (Ev *Ev) Round() *Ev {
	//Round округлить
	Ev.P = math.Round((Ev.P)*1000) / 1000
	Ev.F = math.Round((Ev.F)*1000) / 1000
	Ev.C = math.Round((Ev.C)*1000) / 1000
	Ev.Fb = math.Round((Ev.Fb)*1000) / 1000
	Ev.Cal = math.Round((Ev.Cal)*1000) / 1000
	return Ev
}

func (Ev *Ev) SumEv(Ev2 Ev) Ev {
	Ev.P += Ev2.P
	Ev.F += Ev2.F
	Ev.C += Ev2.C
	Ev.Fb += Ev2.Fb
	Ev.Cal += Ev2.Cal
	Ev3 := *Ev
	return Ev3
}
func (Ev *Ev) DiffEv(Ev2 Ev) Ev {
	Ev.P -= Ev2.P
	Ev.F -= Ev2.F
	Ev.C -= Ev2.C
	Ev.Fb -= Ev2.Fb
	Ev.Cal -= Ev2.Cal
	Ev3 := *Ev
	return Ev3
}

func (Ev *Ev) DivEv() Ev {
	Ev.P /= Ev.W.Weight
	Ev.F /= Ev.W.Weight
	Ev.C /= Ev.W.Weight
	if Ev.Fb != 0 {
		Ev.Fb /= Ev.W.Weight
	}
	Ev.Cal /= Ev.W.Weight
	Ev3 := *Ev
	return Ev3
}

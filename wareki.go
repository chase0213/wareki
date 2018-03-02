package main

type Wareki struct {
	Name  string `json:"name"`
	Yomi  string `json:"yomi"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
}

func (w *Wareki) Seireki() (Seireki, error) {
	return Seireki{}, nil
}

// AD is an alias method of Seireki()
func (w *Wareki) AD() (Seireki, error) {
	return w.Seireki()
}

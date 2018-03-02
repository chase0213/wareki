package wareki

type Wareki struct {
	Name  string `json:"name"`
	Yomi  string `json:"yomi"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
}

func (w *Wareki) Seireki() (Seireki, error) {
	es, err := NewEraSearcher()
	if err != nil {
		return Seireki{}, err
	}

	era, err := es.Search(w.Name)
	seireki := Seireki{
		Year:  era.BeginYear + w.Year - 1,
		Month: w.Month,
		Day:   w.Day,
	}
	return seireki, nil
}

// AD is an alias method of Seireki()
func (w *Wareki) AD() (Seireki, error) {
	return w.Seireki()
}

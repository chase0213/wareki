/*
Package wareki implements a simple converter
for Japanese date format.
*/
package wareki

// Wareki implements converter methods for Japanese date formats
type Wareki struct {
	name  string
	yomi  string
	year  int
	month int
	day   int
}

func (w *Wareki) Name() string {
	return w.name
}

func (w *Wareki) Yomi() string {
	return w.yomi
}

func (w *Wareki) Year() int {
	return w.year
}

func (w *Wareki) Month() int {
	return w.month
}

func (w *Wareki) Day() int {
	return w.day
}

// Seireki converts given date formatted as wareki(和暦) into seireki(西暦)
func (w *Wareki) Seireki() (*Seireki, error) {
	es, err := NewEraSearcher()
	if err != nil {
		return nil, err
	}

	era, err := es.Search(w.name)
	seireki := &Seireki{
		year:  era.BeginYear + w.year - 1,
		month: w.month,
		day:   w.day,
	}
	return seireki, nil
}

// AD is an alias method of Seireki()
func (w *Wareki) AD() (*Seireki, error) {
	return w.Seireki()
}

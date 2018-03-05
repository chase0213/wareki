/*
Package wareki implements a simple converter
for Japanese date format.
*/
package wareki

// Wareki implements converter methods for Japanese date formats
type Wareki interface {
	Name()
	Yomi()
	Year()
	Month()
	Day()
	Seireki()
}

type wareki struct {
	name  string
	yomi  string
	year  int
	month int
	day   int
}

func (w *wareki) Name() string {
	return w.name
}

func (w *wareki) Yomi() string {
	return w.yomi
}

func (w *wareki) Year() int {
	return w.year
}

func (w *wareki) Month() int {
	return w.month
}

func (w *wareki) Day() int {
	return w.day
}

// Seireki converts given date formatted as wareki(和暦) into seireki(西暦)
func (w *wareki) Seireki() (*seireki, error) {
	es, err := NewEraSearcher()
	if err != nil {
		return nil, err
	}

	era, err := es.Search(w.name)
	seireki := &seireki{
		year:  era.BeginYear + w.year - 1,
		month: w.month,
		day:   w.day,
	}
	return seireki, nil
}

// AD is an alias method of Seireki()
func (w *wareki) AD() (*seireki, error) {
	return w.Seireki()
}

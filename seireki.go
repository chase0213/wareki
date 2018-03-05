package wareki

import (
	"errors"
)

// Seireki implements converter methods for Japanese date formats
type Seireki interface {
	Year()
	Month()
	Day()
	Wareki()
}

type seireki struct {
	year  int
	month int
	day   int
}

func (s *seireki) Year() int {
	return s.year
}

func (s *seireki) Month() int {
	return s.month
}

func (s *seireki) Day() int {
	return s.day
}

// Wareki converts given date formatted as seireki(西暦) into wareki(和暦)
func (s *seireki) Wareki() (*wareki, error) {
	eras, err := LoadEras()
	if err != nil {
		return nil, err
	}

	lPtr, rPtr := 0, len(eras)-1
	cPtr := int((rPtr - lPtr) / 2)
	for true {
		center := eras[cPtr]

		//
		// beginning of the first era < target date
		//
		first := eras[0]
		if CompareDate(s.year, s.month, s.day, first.BeginYear, first.BeginMonth, first.BeginDay) < 0 {
			return &wareki{
				name:  "【元号不明】",
				yomi:  "",
				year:  s.year,
				month: s.month,
				day:   s.day,
			}, nil
		}

		//
		// beginning of current era >= target date
		//
		current := eras[rPtr]
		if CompareDate(s.year, s.month, s.day, current.BeginYear, current.BeginMonth, current.BeginDay) >= 0 {
			return &wareki{
				name:  current.Name,
				yomi:  current.Yomi,
				year:  s.year - current.BeginYear + 1,
				month: s.month,
				day:   s.day,
			}, nil
		}

		//
		// beginning of era <= target date < end of era
		//
		if CompareDate(s.year, s.month, s.day, center.BeginYear, center.BeginMonth, center.BeginDay) >= 0 &&
			CompareDate(s.year, s.month, s.day, center.EndYear, center.EndMonth, center.EndDay) <= 0 {
			return &wareki{
				name:  center.Name,
				yomi:  center.Yomi,
				year:  s.year - center.BeginYear + 1,
				month: s.month,
				day:   s.day,
			}, nil
		}

		// if lPtr = rPtr = cPtr and no era hits
		if lPtr >= rPtr {
			return nil, errors.New("invalid date: no era found")
		}

		//
		// beginning of era > target date
		//
		if CompareDate(s.year, s.month, s.day, center.BeginYear, center.BeginMonth, center.BeginDay) == -1 {
			rPtr = cPtr - 1
			cPtr = int((rPtr-lPtr)/2) + lPtr
		}

		//
		// target date <= end of era
		//
		if CompareDate(s.year, s.month, s.day, center.EndYear, center.EndMonth, center.EndDay) >= 0 {
			lPtr = cPtr + 1
			cPtr = int((rPtr-lPtr)/2) + lPtr
		}
	}
	return nil, nil
}

// CompareDate compares two dates, and returns
//  1 if date1 > date2,
// -1 if date1 < date2,
//  0 if date1 = date2
func CompareDate(year1 int, month1 int, day1 int, year2 int, month2 int, day2 int) int {
	if year1 > year2 {
		return 1
	}

	if year1 < year2 {
		return -1
	}

	if month1 > month2 {
		return 1
	}

	if month1 < month2 {
		return -1
	}

	if day1 > day2 {
		return 1
	}

	if day1 < day2 {
		return -1
	}

	return 0
}

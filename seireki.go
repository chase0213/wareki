package main

import (
	"errors"
)

type Seireki struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (s *Seireki) Wareki() (Wareki, error) {
	lPtr, rPtr := 0, len(eraDict)-1
	cPtr := int((rPtr - lPtr) / 2)
	for true {
		center := eraDict[cPtr]

		//
		// beginning of current era >= target date
		//
		current := eraDict[rPtr]
		if CompareDate(s.Year, s.Month, s.Day, current.BeginYear, current.BeginMonth, current.BeginDay) >= 0 {
			return Wareki{
				Name:  current.Name,
				Yomi:  current.Yomi,
				Year:  s.Year - current.BeginYear + 1,
				Month: s.Month,
				Day:   s.Day,
			}, nil
		}

		//
		// beginning of era <= target date < end of era
		//
		if CompareDate(s.Year, s.Month, s.Day, center.BeginYear, center.BeginMonth, center.BeginDay) >= 0 &&
			CompareDate(s.Year, s.Month, s.Day, center.EndYear, center.EndMonth, center.EndDay) <= 0 {
			return Wareki{
				Name:  center.Name,
				Yomi:  center.Yomi,
				Year:  s.Year - center.BeginYear + 1,
				Month: s.Month,
				Day:   s.Day,
			}, nil
		}

		// if lPtr = rPtr = cPtr and no era hits
		if lPtr >= rPtr {
			return Wareki{}, errors.New("invalid date: no era found")
		}

		//
		// beginning of era > target date
		//
		if CompareDate(s.Year, s.Month, s.Day, center.BeginYear, center.BeginMonth, center.BeginDay) == -1 {
			rPtr = cPtr - 1
			cPtr = int((rPtr-lPtr)/2) + lPtr
		}

		//
		// target date <= end of era
		//
		if CompareDate(s.Year, s.Month, s.Day, center.EndYear, center.EndMonth, center.EndDay) >= 0 {
			lPtr = cPtr + 1
			cPtr = int((rPtr-lPtr)/2) + lPtr
		}
	}
	return Wareki{}, nil
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

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func ParseWarekiString(dateStr string) (Wareki, error) {
	dateWithoutNumber := regexp.MustCompile("一|二|三|四|五|六|七|八|九|十|百|千|零|元|[0-9０-９]+").Split(dateStr, 2)
	if len(dateWithoutNumber) < 2 {
		return Wareki{}, fmt.Errorf("Invalid date format")
	}

	es, err := NewEraSearcher()
	if err != nil {
		return Wareki{}, err
	}

	nengo := dateWithoutNumber[0]
	era, err := es.Search(nengo)
	if err != nil {
		return Wareki{}, err
	}

	dateSlice := regexp.MustCompile(era.Name+"|/|,|-|\\s|年|月|日").Split(dateStr, 5)
	year, _ := numberize(dateSlice[1])
	month, _ := numberize(dateSlice[2])
	day, _ := numberize(dateSlice[3])

	wareki := Wareki{
		Name:  era.Name,
		Yomi:  era.Yomi,
		Year:  year,
		Month: month,
		Day:   day,
	}

	return wareki, nil
}

func ParseSeirekiString(dateStr string) (Seireki, error) {
	dateSlice := regexp.MustCompile("/|,|-|\\s|年|月|日").Split(dateStr, 4)
	if len(dateSlice) < 3 {
		err := fmt.Errorf("Invalid date format: date should be separated by '/', ',', '-', or whitespaces")
		return Seireki{}, err
	}

	year, _ := strconv.Atoi(dateSlice[0])
	month, _ := strconv.Atoi(dateSlice[1])
	day, _ := strconv.Atoi(dateSlice[2])
	return Seireki{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
}

func numberize(numberStr string) (int, error) {
	// 変換できる場合は単純に変換して返す
	if num, err := strconv.Atoi(numberStr); err == nil {
		return num, nil
	}

	// 漢数字の場合
	num := 0
	n := 1
	isSubunit := false

	for _, l := range numberStr {
		// 数字
		res, err := number(l)
		if err == nil {
			if n >= 10 {
				num += n
			}
			n = res

			isSubunit = false
			continue
		}

		// 補助単位
		res, err = subunit(l)
		if err == nil {
			// 補助単位が2回続いた場合は、単位をリセット
			if isSubunit {
				num += n
				n = 1
			}

			n *= res
			isSubunit = true
			continue
		}

		// 単位
		res, err = unit(l)
		if err == nil {
			n *= res
			num += n
			n = 1

			isSubunit = false
			continue
		}

		return 0, fmt.Errorf("Invalid kanji-value: %s", numberStr)
	}

	// 一の位が残ることがあるので、それを最後に足す
	if n != 0 {
		num += n
	}

	return num, nil
}

func number(char rune) (int, error) {
	switch char {
	case '零':
		return 0, nil
	case '元', '一':
		return 1, nil
	case '二':
		return 2, nil
	case '三':
		return 3, nil
	case '四':
		return 4, nil
	case '五':
		return 5, nil
	case '六':
		return 6, nil
	case '七':
		return 7, nil
	case '八':
		return 8, nil
	case '九':
		return 9, nil
	default:
		return 0, fmt.Errorf("Invalid Kanji-number: %s", string(char))
	}
}

func unit(char rune) (int, error) {
	switch char {
	case '億':
		return 100000000, nil
	case '万':
		return 10000, nil
	case '千':
		return 1000, nil
	case '百':
		return 100, nil
	case '十':
		return 10, nil
	default:
		return 0, fmt.Errorf("Invalid Kanji-unit: %s", string(char))
	}
}

func subunit(char rune) (int, error) {
	switch char {
	case '千':
		return 1000, nil
	case '百':
		return 100, nil
	case '十':
		return 10, nil
	default:
		return 0, fmt.Errorf("Invalid Kanji-subunit: %s", string(char))
	}
}

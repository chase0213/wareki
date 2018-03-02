package wareki

import "testing"

func TestParseWarekiString(t *testing.T) {
	var text string

	text = "昭和三十六年五月二十三日"
	wareki, err := ParseWarekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if wareki.Name != "昭和" || wareki.Year != 36 || wareki.Month != 5 || wareki.Day != 23 {
		t.Errorf("expected: 昭和36年5月23日\nactual: %s%d年%d月%d日\n", wareki.Name, wareki.Year, wareki.Month, wareki.Day)
	}

	text = "平成元年二月十三日"
	wareki, err = ParseWarekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if wareki.Name != "平成" || wareki.Year != 1 || wareki.Month != 2 || wareki.Day != 13 {
		t.Errorf("expected: 平成1年2月13日\nactual: %s%d年%d月%d日\n", wareki.Name, wareki.Year, wareki.Month, wareki.Day)
	}

	text = "平成1年2月13日"
	wareki, err = ParseWarekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if wareki.Name != "平成" || wareki.Year != 1 || wareki.Month != 2 || wareki.Day != 13 {
		t.Errorf("expected: 平成1年2月13日\nactual: %s%d年%d月%d日\n", wareki.Name, wareki.Year, wareki.Month, wareki.Day)
	}

	text = "新号元年二月十三日"
	wareki, err = ParseWarekiString(text)
	if err == nil {
		t.Errorf("parser should return error if no 元号 found: %s", text)
	}
}

func TestParseSeirekiString(t *testing.T) {
	var text string

	text = "2018/02/13"
	seireki, err := ParseSeirekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if seireki.Year != 2018 || seireki.Month != 2 || seireki.Day != 13 {
		t.Errorf("expected: 2018/02/13\nactual: %d/%d/%d\n", seireki.Year, seireki.Month, seireki.Day)
	}

	text = "1989-12-8"
	seireki, err = ParseSeirekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if seireki.Year != 1989 || seireki.Month != 12 || seireki.Day != 8 {
		t.Errorf("expected: 1989/02/13\nactual: %d/%d/%d\n", seireki.Year, seireki.Month, seireki.Day)
	}

	text = "2020-120-813"
	seireki, err = ParseSeirekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if seireki.Year != 2020 || seireki.Month != 120 || seireki.Day != 813 {
		t.Errorf("expected: 2020/120/813\nactual: %d/%d/%d\n", seireki.Year, seireki.Month, seireki.Day)
	}
}

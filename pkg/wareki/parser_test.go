package wareki

import "testing"

func TestParseWarekiString(t *testing.T) {
	var text string

	text = "昭和三十六年五月二十三日"
	wareki, err := ParseWarekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if wareki.name != "昭和" || wareki.year != 36 || wareki.month != 5 || wareki.day != 23 {
		t.Errorf("expected: 昭和36年5月23日\nactual: %s%d年%d月%d日\n", wareki.name, wareki.year, wareki.month, wareki.day)
	}

	text = "平成元年二月十三日"
	wareki, err = ParseWarekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if wareki.name != "平成" || wareki.year != 1 || wareki.month != 2 || wareki.day != 13 {
		t.Errorf("expected: 平成1年2月13日\nactual: %s%d年%d月%d日\n", wareki.name, wareki.year, wareki.month, wareki.day)
	}

	text = "平成1年2月13日"
	wareki, err = ParseWarekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if wareki.name != "平成" || wareki.year != 1 || wareki.month != 2 || wareki.day != 13 {
		t.Errorf("expected: 平成1年2月13日\nactual: %s%d年%d月%d日\n", wareki.name, wareki.year, wareki.month, wareki.day)
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

	if seireki.year != 2018 || seireki.month != 2 || seireki.day != 13 {
		t.Errorf("expected: 2018/02/13\nactual: %d/%d/%d\n", seireki.year, seireki.month, seireki.day)
	}

	text = "1989-12-8"
	seireki, err = ParseSeirekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if seireki.year != 1989 || seireki.month != 12 || seireki.day != 8 {
		t.Errorf("expected: 1989/02/13\nactual: %d/%d/%d\n", seireki.year, seireki.month, seireki.day)
	}

	text = "2020-120-813"
	seireki, err = ParseSeirekiString(text)
	if err != nil {
		t.Errorf("cannot parse %s", text)
	}

	if seireki.year != 2020 || seireki.month != 120 || seireki.day != 813 {
		t.Errorf("expected: 2020/120/813\nactual: %d/%d/%d\n", seireki.year, seireki.month, seireki.day)
	}
}

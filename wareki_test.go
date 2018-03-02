package wareki

import "testing"

func TestWarekiToSeireki(t *testing.T) {
	var wareki Wareki

	wareki.Name = "平成"
	wareki.Year = 30
	wareki.Month = 2
	wareki.Day = 13

	seireki, err := wareki.Seireki()
	if err != nil {
		t.Errorf("%s%d年%d月%d日 should be parsed successfully", wareki.Name, wareki.Year, wareki.Month, wareki.Day)
	}

	if seireki.Year != 2018 || seireki.Month != 2 || seireki.Day != 13 {
		t.Errorf("expected\t2018/2/13\ngot\t%d/%d/%d", seireki.Year, seireki.Month, seireki.Day)
	}

	wareki.Name = "昭和"
	wareki.Year = 64
	wareki.Month = 1
	wareki.Day = 7

	seireki, err = wareki.Seireki()
	if err != nil {
		t.Errorf("%s%d年%d月%d日 should be parsed successfully", wareki.Name, wareki.Year, wareki.Month, wareki.Day)
	}

	if seireki.Year != 1989 || seireki.Month != 1 || seireki.Day != 7 {
		t.Errorf("expected\t1989/1/7\ngot\t%d/%d/%d", seireki.Year, seireki.Month, seireki.Day)
	}

	wareki.Name = "存在しない元号"
	wareki.Year = 123
	wareki.Month = 2
	wareki.Day = 13

	seireki, err = wareki.Seireki()
	if err != nil {
		t.Errorf("%s%d年%d月%d日 should be parsed successfully", wareki.Name, wareki.Year, wareki.Month, wareki.Day)
	}

	if seireki.Year != 122 || seireki.Month != 2 || seireki.Day != 13 {
		t.Errorf("expected\122/2/13\ngot\t%d/%d/%d", seireki.Year, seireki.Month, seireki.Day)
	}
}
